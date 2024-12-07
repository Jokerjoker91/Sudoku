document
  .getElementById("generate-btn")
  .addEventListener("click", async function () {
    try {
      const apiUrl =
        location.hostname === "localhost" || location.hostname === "127.0.0.1"
          ? "http://localhost:8080/generate?difficulty=5" // Local URL for development
          : "https://sudoku-yy8u.onrender.com/generate"; // GitHub Pages URL for production

      const response = await fetch(apiUrl); // Fetch the grid from the appropriate URL
      const sudokuGrid = await response.json();

      const sudokuContainer = document.getElementById("sudoku-container");
      sudokuContainer.innerHTML = ""; // Clear any existing grid

      sudokuGrid.forEach((row, rowIndex) => {
        row.forEach((cell, colIndex) => {
          const cellElement = document.createElement("div");
          cellElement.classList.add("sudoku-cell");
          cellElement.contentEditable = cell === 0; // Make cell editable if empty
          cellElement.textContent = cell !== 0 ? cell : ""; // Show empty cell if value is 0

          if (rowIndex % 3 === 2) cellElement.classList.add("highlight-row");
          if (colIndex % 3 === 2) cellElement.classList.add("highlight-column");

          // Add the .initial class for cells that are pre-filled (non-null values)
          if (cell !== 0) {
            cellElement.classList.add("initial");
          }

          // Event listener for input change
          cellElement.addEventListener("input", async (e) => {
            const value = parseInt(e.target.textContent, 10);
            if (isNaN(value)) {
              return; // Ignore if the input is not a number
            }

            const apiUrl =
              location.hostname === "localhost" ||
              location.hostname === "127.0.0.1"
                ? "http://localhost:8080/validate" // Local URL for development
                : "https://sudoku-yy8u.onrender.com/validate"; // GitHub Pages URL for production

            // Send the number to the backend for validation
            const response = await fetch(apiUrl, {
              method: "POST",
              headers: { "Content-Type": "application/json" },
              body: JSON.stringify({
                grid: sudokuGrid,
                row: rowIndex,
                col: colIndex,
                number: value,
              }),
            });
            const result = await response.json();

            // Change the color based on whether the number is valid
            if (result.valid) {
              e.target.style.color = "blue";
            } else {
              e.target.style.color = "red";
            }

            // Update the grid with the new value
            sudokuGrid[rowIndex][colIndex] = value;
          });

          sudokuContainer.appendChild(cellElement);
        });
      });

      // Enable the "Solve Sudoku" button once the grid is generated
      document.getElementById("solve-btn").disabled = false;
    } catch (error) {
      console.error("Error generating Sudoku:", error);
    }
  });

document
  .getElementById("solve-btn")
  .addEventListener("click", async function () {
    const sudokuContainer = document.getElementById("sudoku-container");
    const cells = Array.from(
      sudokuContainer.getElementsByClassName("sudoku-cell")
    );
    const sudokuGrid = [];

    // Build the grid from the current input
    for (let i = 0; i < 9; i++) {
      const row = [];
      for (let j = 0; j < 9; j++) {
        const cellValue = parseInt(cells[i * 9 + j].textContent, 10);
        row.push(isNaN(cellValue) ? 0 : cellValue);
      }
      sudokuGrid.push(row);
    }

    try {
      const apiUrl =
        location.hostname === "localhost" || location.hostname === "127.0.0.1"
          ? "http://localhost:8080/solve" // Local URL for development
          : "https://sudoku-yy8u.onrender.com/solve"; // GitHub Pages URL for production

      const response = await fetch(apiUrl, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ grid: sudokuGrid }),
      });
      const result = await response.json();

      // Check if the puzzle is solvable
      if (result.solvable) {
        // Update the grid with the solved values
        const solution = result.solution;
        cells.forEach((cell, index) => {
          const row = Math.floor(index / 9);
          const col = index % 9;
          const solutionValue = solution[row][col];
          cell.textContent = solutionValue !== 0 ? solutionValue : "";
        });
      } else {
        alert("This Sudoku puzzle is not solvable.");
      }
    } catch (error) {
      console.error("Error solving Sudoku:", error);
    }
  });
