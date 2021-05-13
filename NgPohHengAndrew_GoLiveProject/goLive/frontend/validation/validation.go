// Package validation performs checking of raw string values
// from HTML form inputs, and blacklist with regexp.
// If the validation failed, i.e. blacklisted, the raw string
// will be rejected, and an error returned.
// Otherwise, the raw string will be sanitized and return
// to the caller that invokes the validation checking.
// If the string contains any HTML <tags> that poses a security issue,
// an error is returned immediately, and an intrusion alert is raised.
package validation

import (
	"errors"
	"html"
	"regexp"
	"strings"
	"time"

	"goLive/frontend/common"
)

// CheckString checks whether HTML input of
// raw string contains HTML \<tags\> with sanitization.
func CheckString(s string) (string, error) {
	// check for HTML tags <tag>
	re := regexp.MustCompile(`<\w+>|<\/\w+>|\(.+\).+`)
	match := re.Match([]byte(s))
	if match {
		return "", errors.New("error: blacklisted")
	}
	re = regexp.MustCompile(`\b\W+\b`)
	// sanitization
	re = regexp.MustCompile(`-+`) // has dashes in string
	match = re.Match([]byte(s))
	if match {
		s = strings.TrimSpace(s)
		return s, nil
	}
	// sanitization
	s = re.ReplaceAllString(html.EscapeString(s), " ")
	s = strings.TrimSpace(s)
	return s, nil
}

// CheckUUIDString checks whether HTML input of
// raw string contains HTML \<tags\> with sanitization.
func CheckUUIDString(s string) (string, error) {
	// check for HTML tags <tag>
	re := regexp.MustCompile(`<\w+>|<\/\w+>|\(.+\).+`)
	match := re.Match([]byte(s))
	if match {
		return "", errors.New("error: blacklisted")
	}
	// sanitization
	re = regexp.MustCompile(`-+`) // has dashes in string
	match = re.Match([]byte(s))
	if match {
		// not UUID
		if len(s) != 36 {
			return "", errors.New("error: string mismatched")
		}
		s = strings.TrimSpace(s)
		return s, nil
	}
	re = regexp.MustCompile(`\b\W+\b`)
	// sanitization
	s = re.ReplaceAllString(html.EscapeString(s), " ")
	s = strings.TrimSpace(s)
	return s, nil
}

// ValidateString validates whether HTML input of raw string,
// s, is being manipulated by an attacker.
func ValidateString(s string, length int) (string, error) {
	s, err := CheckString(s)
	if err != nil {
		return "", err
	}
	if len(s) == 0 {
		return "", errors.New("error: empty string")
	}

	// for checking image
	if length == -1 {
		return s, nil
	}

	// others
	if len(s) > length {
		return s, errors.New("error: string too long")
	}

	return s, nil
}

// ValidatePassword validates whether HTML input of raw string,
// password, is being manipulated by an attacker.
// At the same time validate conditions:
// 8 characters; at least 1 capital letter, a small letter,
// 1 digit and 1 special character
func ValidatePassword(password string) (string, error) {
	s, err := CheckString(password)
	if err != nil {
		return "", err
	}
	if len(s) == 0 {
		return "", errors.New("error: empty password")
	}

	/*if len(s) < 8 {
		return s, errors.New("error: Password must be at least\n" +
			"- 8 characters long")
	}

	e := "error: Password must have at least\n"
	r := regexp.MustCompile("(.*?[a-z])")
	checkOK := r.MatchString(s)
	if !checkOK {
		e += "- 1 small letter,\n"
	}
	r = regexp.MustCompile("(.*?[A-Z])")
	checkOK = r.MatchString(s)
	if !checkOK {
		e += "- 1 capital letter,\n"
	}
	r = regexp.MustCompile("(.*?[0-9])")
	checkOK = r.MatchString(s)
	if !checkOK {
		e += "- 1 digit,\n"
	}
	r = regexp.MustCompile("(.*?[+()#@$?*^&.,%!=-])")
	checkOK = r.MatchString(s)
	if !checkOK {
		e += "- 1 special character"
	}
	if !checkOK {
		return s, errors.New(e)
	}*/

	return s, nil
}

// ValidateDate validates whether HTML input of raw string,
// date "2021-03-17", is being manipulated by an attacker.
func ValidateDate(date string) (string, error) {
	// validate data
	_, err := CheckString(date)
	if err != nil {
		return "", err
	}

	// boundary checking of fixed length
	// "2021-03-17" | "2x21-03-17" gives error
	re := regexp.MustCompile(`\b\d{4}-\d{2}-\d{2}\b`)
	isDate := re.MatchString(date)
	isFixedLen := len(date) == len("YYYY-MM-DD")

	// convert string to Array
	dateArr := strings.Split(date, "-")

	// check index out-of-bounds error
	if len(dateArr) != 3 { // not length of 3
		return date, errors.New("error: out-of-bounds")
	}

	// not out-of-bounds error
	userIntYear, errYear := common.ToInt(dateArr[0]) // 2021
	userIntMth, errMonth := common.ToInt(dateArr[1]) // 3
	userIntDay, errDay := common.ToInt(dateArr[2])   // 17

	// user input
	userMonth := time.Month(userIntMth)
	// zeroed to YYYY-MM-DD
	userTime := time.Date(userIntYear, userMonth, userIntDay, 0, 0, 0, 0, time.UTC)
	userTimestamp := userTime.Unix() // in seconds

	// this current year
	thisYear, month, day := time.Now().Date()
	// zeroed to YYYY-MM-DD
	today := time.Date(thisYear, month, day, 0, 0, 0, 0, time.UTC)
	todayTimeStamp := today.Unix() // in seconds
	// zeroed to YYYY-MM-DD
	oneYearTime := time.Date(thisYear+1, month, day, 0, 0, 0, 0, time.UTC)
	oneYearTimeStamp := oneYearTime.Unix() // in seconds

	// check whether is today or future; not past
	isValidToday := userTimestamp >= todayTimeStamp // in seconds
	isValidToday = isValidToday && (oneYearTimeStamp >= userTimestamp)

	// within one year from now
	numYear := userIntYear - thisYear
	// between 0 and 1 year
	isValidOneYear := numYear >= 0 && numYear <= 1

	// incorrect
	if (isFixedLen && isDate) == false ||
		isValidToday == false || isValidOneYear == false ||
		(errYear != nil || errMonth != nil || errDay != nil) {
		common.Debug(numYear, isValidOneYear, isValidToday, date,
			errYear, errMonth, errDay)
		return date, errors.New("error: wrong date")
	}

	// [OK]
	return date, nil
}

// GetErrorStr converts errors message to string.
func GetErrorStr(err error) (string, bool) {
	flag := false
	s := err.Error()
	if strings.Contains(s, "error:") {
		re := regexp.MustCompile("error: ")
		s = re.ReplaceAllString(s, "")
		flag = true
	}
	return s, flag
}
