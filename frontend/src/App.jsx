import { useState } from 'react'
import './App.css'

// const URL = import.meta.env.VITE_API_URL

function App() {
  const [boardSize, setBoardSize] = useState(15);
  const [playerTurn, setPlayerTurn] = useState("X");
  const [board, setBoard] = useState(createBoard(boardSize));
  const [isStarted, setIsStarted] = useState(false);
  const [isOver, setIsOver] = useState(false);
  const [isBlank, setIsBlank] = useState(true);
  const [isThinking, setIsThinking] = useState(false);

  function createBoard(size) {
    return Array.from({ length: size }, () => Array.from({ length: size }, () => ""))
  }

  const handlePlayerSelect = (e) => {
    setPlayerTurn(e.target.value);
  };

  function handleClick(row, col) {
    if (board[row][col] === "") {
      const updatedBoard = [...board];
      updatedBoard[row][col] = playerTurn;
      setBoard(updatedBoard);
      setIsThinking(true);

      fetch(`/api/move`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          'Access-Control-Allow-Origin': '*'
        },
        body: JSON.stringify({ "Row": row + 1, "Col": col + 1 })
      })
        .then(response => response.json())
        .then(data => {
          console.log("data: ", data)

          setIsThinking(false);
          if (data.GameOver) {
            setIsOver(true);
            console.log("Game Over")
          }

          const row = data.Move.Row - 1;
          const col = data.Move.Col - 1;
          const updatedBoard = [...board];
          updatedBoard[row][col] = playerTurn === "X" ? "O" : "X";
          setBoard(updatedBoard);

        })
        .catch(error => {
          console.log(error)
        });
    }
  }

  function handleBoardSize(e) {
    setBoardSize(parseInt(e.target.value))
    setBoard(createBoard(parseInt(e.target.value)))
  }

  function handleStart() {
    setIsBlank(false);
    setIsOver(false);
    setPlayerTurn("X");
    setIsThinking(false);
    fetch(`/api/init`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Access-Control-Allow-Origin': '*'
      },
      body: JSON.stringify({ "BoardSize": boardSize, "Player": playerTurn })
    })
      .then(response => response.json())
      .then(data => {
        if (playerTurn === "X") {
          const row = data.FirstMove.Row - 1;
          const col = data.FirstMove.Col - 1;
          const updatedBoard = [...board];
          updatedBoard[row][col] = playerTurn
          setPlayerTurn("O");
          setBoard(updatedBoard);
          setIsThinking(false);
        }
        console.log("data: ", data)
      })
      .catch(error => {
        console.log(error)
      });
    setIsStarted(true);
  }

  function handleNewGame() {
    setIsBlank(true);
    setBoard(createBoard(boardSize));
    setIsStarted(false);
  }

  return (
    <>
      <input
        type="number"
        value={boardSize}
        onChange={(e) => handleBoardSize(e)}
      />
      <div className="board">
        {Array.from({ length: boardSize }, (_, row) => (
          <div key={row} className="row">
            {Array.from({ length: boardSize }, (_, col) => (
              <button
                key={`${row}-${col}`}
                className="cell"
                onClick={() => handleClick(row, col)}
                disabled={!isStarted || isOver || isThinking}
              >
                {board[row][col] === "X" && <span>X</span>}
                {board[row][col] === "O" && <span>O</span>}
              </button>
            ))}
          </div>
        ))}
      </div>
      <div><span>AI Player: </span>
        <select value={playerTurn} onChange={handlePlayerSelect} disabled={isStarted}>
          <option value="X">X</option>
          <option value="O">O</option>
        </select>
        <button onClick={handleNewGame}>New game</button>
        <button onClick={handleStart} disabled={!isBlank}>Start Game</button>
      </div>
      {isOver && <div>Game Over</div>}
    </>
  );
}

export default App
