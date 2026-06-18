import { useState, useEffect } from "react";
import "./App.css";

function App() {
  const [topic, setTopic] = useState("");
  const [currentTopic, setCurrentTopic] = useState("-");
  const [researchCount, setResearchCount] = useState(0);

  useEffect(() => {
    const interval = setInterval(async () => {
      try {
        const response = await fetch(
          "http://localhost:8080/research/status"
        );

        const data = await response.json();

        setResearchCount(data.count || 0);

        if (data.currentTopic) {
          setCurrentTopic(data.currentTopic);
        }
      } catch (err) {
        console.error(err);
      }
    }, 2000);

    return () => clearInterval(interval);
  }, []);

  const startResearch = async () => {
    if (!topic.trim()) return;

    try {
      const response = await fetch(
        "http://localhost:8080/research/start",
        {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({
            topic: topic,
          }),
        }
      );

      if (!response.ok) {
        throw new Error("Failed to start research");
      }

      setCurrentTopic(topic);

      console.log("Research started:", topic);
    } catch (err) {
      console.error(err);
    }
  };

  const stopResearch = async () => {
    try {
      await fetch(
        "http://localhost:8080/research/stop",
        {
          method: "POST",
        }
      );

      console.log("Research stopped");
    } catch (err) {
      console.error(err);
    }
  };

return (
  <div className="container">
    <h1>Autonomous Researcher</h1>

    <div className="card">
      <input
        type="text"
        placeholder="Enter research topic..."
        value={topic}
        onChange={(e) =>
          setTopic(e.target.value)
        }
      />

      <div className="buttons">
        <button onClick={startResearch}>
          Start Research
        </button>

        <button
          className="stop-btn"
          onClick={stopResearch}
        >
          Stop Research
        </button>
      </div>
    </div>

    <div className="stats">
      <div className="stat-card">
        <h3>Current Topic</h3>
        <p>{currentTopic}</p>
      </div>

      <div className="stat-card">
        <h3>Research Records</h3>
        <p>{researchCount}</p>
      </div>
    </div>
  </div>
);
}

export default App;