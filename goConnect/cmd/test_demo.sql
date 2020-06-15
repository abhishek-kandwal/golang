-- phpMyAdmin SQL Dump
-- version 4.6.6deb5
-- https://www.phpmyadmin.net/
--
-- Host: localhost:3306
-- Generation Time: Jun 10, 2020 at 04:49 PM
-- Server version: 5.7.30-0ubuntu0.18.04.1
-- PHP Version: 7.2.24-0ubuntu0.18.04.3

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `test_demo`
--

-- --------------------------------------------------------

--
-- Table structure for table `weathermap`
--

CREATE TABLE `weathermap` (
  `id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `visibility` varchar(255) NOT NULL,
  `base` varchar(255) NOT NULL,
  `dt` varchar(255) NOT NULL,
  `cond` varchar(255) NOT NULL,
  `lon` varchar(255) NOT NULL,
  `lat` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `weathermap`
--

INSERT INTO `weathermap` (`id`, `name`, `visibility`, `base`, `dt`, `cond`, `lon`, `lat`) VALUES
(2, 'america', 'visiable', 'base', 'dds', 'coons', 'lll', '12.34'),
(3, 'newyork', 'visiable', 'base', 'dds', 'coons', 'lll', '12.34');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `weathermap`
--
ALTER TABLE `weathermap`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `weathermap`
--
ALTER TABLE `weathermap`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
