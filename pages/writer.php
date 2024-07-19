<?php
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

foreach( $data as $row ) {
    $query .= 
                "INSERT INTO TestTable ( root, src, parsed_time, date) VALUES 
                ('".$row["path"]."', '".$row["data[3]"]."', 
                '".$row["elapsedTime"]."', now(); ";
}

$conn->close();
?> 