-- phpMyAdmin SQL Dump
-- version 4.9.5
-- https://www.phpmyadmin.net/
--
-- Host: localhost:3306
-- Generation Time: May 04, 2021 at 07:04 PM
-- Server version: 5.6.41-84.1
-- PHP Version: 7.3.27

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `sacmcoms_auction`
--

-- --------------------------------------------------------

--
-- Table structure for table `salesItems`
--

CREATE TABLE `salesItems` (
  `itemID` bigint(20) UNSIGNED NOT NULL,
  `userID` int(11) NOT NULL,
  `itemName` varchar(50) COLLATE utf8_unicode_ci NOT NULL,
  `itemDesc` varchar(100) COLLATE utf8_unicode_ci NOT NULL,
  `bidClosebyOwner` tinyint(1) NOT NULL,
  `bidCloseDate` datetime NOT NULL,
  `bidIncrement` int(4) NOT NULL,
  `basePrice` int(6) NOT NULL,
  `displayItem` tinyint(1) NOT NULL,
  `itemImage` blob NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci COMMENT='This table records items put up for auction.  ';

--
-- Indexes for dumped tables
--

--
-- Indexes for table `salesItems`
--
ALTER TABLE `salesItems`
  ADD PRIMARY KEY (`itemID`),
  ADD UNIQUE KEY `itemID` (`itemID`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `salesItems`
--
ALTER TABLE `salesItems`
  MODIFY `itemID` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
