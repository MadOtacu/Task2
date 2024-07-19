<?php
ini_set('display_errors', 1);
error_reporting(E_ALL);

$servername = "localhost";
$username = "root";
$password = "password";
$dbname = "TestDB";

// Create connection
$conn = new mysqli($servername, $username, $password, $dbname);
// Check connection
if ($conn->connect_error) {
  die("Connection failed: " . $conn->connect_error);
}

$data = json_decode(file_get_contents('php://input'), true);

$date = date('Y-m-d');
echo implode(" ", $data);

$root = $data['root'];
$size = $data['size'];
$time_spent = $data['parsedTime'];

$query = $conn->prepare("INSERT INTO TestTable (root, src, parsed_time, date) VALUES (?, ?, ?, NOW())");
$query->bind_param("sii", $root, $size, $time_spent);

$query->execute();

$conn->close();
?> 