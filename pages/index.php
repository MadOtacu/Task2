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

$sql = "SELECT id, root, src, parsed_time, date FROM TestTable";
$result = $conn->query($sql);

if ($result->num_rows > 0) {
  // output data of each row
  while($row = $result->fetch_assoc()) {
    echo "id: " . $row["id"]. " - root: " . $row["root"]. " - size: " . $row["src"]. " - parsed_time: " . $row["parsed_time"]. " - date: " . $row["date"]. "<br>";
  }
} else {
  echo "0 results";
}
$conn->close();
?> 