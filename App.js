import React, { useEffect, useState } from "react";
import axios from "axios";

function App() {
  const [sensorData, setSensorData] = useState([]);

  useEffect(() => {
    fetchData();
  }, []);

  const fetchData = async () => {
    try {
      const response = await axios.get("http://localhost:8080/get-data");
      setSensorData(response.data);
    } catch (error) {
      console.error("Error fetching sensor data", error);
    }
  };

  return (
    <div>
      <h1>Environmental Monitoring Dashboard</h1>
      <table border="1">
        <thead>
          <tr>
            <th>Timestamp</th>
            <th>Temperature (°C)</th>
            <th>Light (Lux)</th>
            <th>Humidity (%)</th>
          </tr>
        </thead>
        <tbody>
          {sensorData.map((data, index) => (
            <tr key={index}>
              <td>{data.timestamp}</td>
              <td>{data.temperature}</td>
              <td>{data.light}</td>
              <td>{data.humidity}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}

export default App;
