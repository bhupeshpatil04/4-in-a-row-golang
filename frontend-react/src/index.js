
import React, { useEffect, useState, useRef } from "react";
import ReactDOM from "react-dom/client";

const CELL = 60;

function App() {
  const [board, setBoard] = useState(
    Array(6).fill(0).map(() => Array(7).fill(0))
  );
  const [status, setStatus] = useState("Your Turn");
  const ws = useRef(null);

  useEffect(() => {
    ws.current = new WebSocket("ws://localhost:8080/ws");

    ws.current.onmessage = (e) => {
      const game = JSON.parse(e.data);
      setBoard(game.Board);

      if (game.Winner === 1) setStatus("🔴 Player 1 Wins!");
      else if (game.Winner === 2) setStatus("🟡 Player 2 Wins!");
      else setStatus("Your Turn");
    };

    return () => ws.current.close();
  }, []);

  const drop = (c) => {
    ws.current.send(JSON.stringify(c));
  };

  return (
    <div style={styles.page}>
      <h1 style={styles.title}>4 in a Row</h1>

      <div style={styles.card}>
        <div style={styles.status}>{status}</div>

        <div style={styles.board}>
          {board.map((row, r) =>
            row.map((cell, c) => (
              <div
                key={`${r}-${c}`}
                onClick={() => drop(c)}
                style={{
                  ...styles.cell,
                  background:
                    cell === 1 ? "#ef4444" :
                    cell === 2 ? "#facc15" :
                    "#e5e7eb"
                }}
              />
            ))
          )}
        </div>
      </div>

      <p style={styles.footer}>
        GoLang Backend • WebSockets • Real-time Gameplay
      </p>
    </div>
  );
}

const styles = {
  page: {
    minHeight: "100vh",
    background: "#0f172a",
    color: "white",
    display: "flex",
    flexDirection: "column",
    alignItems: "center",
    justifyContent: "center",
    fontFamily: "Inter, system-ui"
  },
  title: {
    marginBottom: 20,
    fontSize: 32,
    fontWeight: 700
  },
  card: {
    background: "#020617",
    padding: 20,
    borderRadius: 16,
    boxShadow: "0 20px 40px rgba(0,0,0,0.4)"
  },
  status: {
    textAlign: "center",
    marginBottom: 12,
    fontSize: 18
  },
  board: {
    display: "grid",
    gridTemplateColumns: `repeat(7, ${CELL}px)`,
    gap: 6
  },
  cell: {
    width: CELL,
    height: CELL,
    borderRadius: "50%",
    cursor: "pointer",
    transition: "transform 0.15s",
  },
  footer: {
    marginTop: 16,
    fontSize: 12,
    opacity: 0.7
  }
};

const root = ReactDOM.createRoot(document.getElementById("root"));
root.render(<App />);

