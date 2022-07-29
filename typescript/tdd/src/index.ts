
interface Grid {
    width: number;
    height: number;
    fields: string[][];
    setFlags: () => string[][];
}

export const getGrid = (width: number, height: number, bombAmount: number): Grid => {

    let bombsPlanted = 0;
    const fields: string[][] = new Array(width).fill("0").map(() => {

        const row = new Array(height).fill("0");

        row.forEach((_, index) => {
           if (bombsPlanted < bombAmount) {
               row[index] = "X";
               bombsPlanted += 1;
           }
       });

        return row;
    });

    fields.forEach((row, rowIndex) => {

        row.forEach((field, fieldIndex) => {
                    if (fields [rowIndex][fieldIndex] !== "X" ) {
                        fields [rowIndex][fieldIndex] = getAmountOfNeighbourBombs( fields, rowIndex, fieldIndex ).toString();
                    }
        });
    });
/*
    fields.forEach((row, rowIndex) => {

        row.forEach((field, fieldIndex) => {
            const siblings = [];
            siblings.push(row[fieldIndex - 1]);
            siblings.push(row[fieldIndex + 1]);
            siblings.push(fields[rowIndex - 1][fieldIndex - 1]);
            siblings.push(fields[rowIndex - 1][fieldIndex]);
            siblings.push(fields[rowIndex - 1][fieldIndex + 1]);
            siblings.push(fields[rowIndex + 1][fieldIndex - 1]);
            siblings.push(fields[rowIndex + 1][fieldIndex]);
            siblings.push(fields[rowIndex + 1][fieldIndex + 1]);

            row[fieldIndex] = siblings.filter((f) => f === "X").length.toString();

        });
    });
*/
    return {
        width,
        height,
        fields,
        setFlags: () => [],
    };
};

export function getAmountOfNeighbourBombs(myGrid: string[][] , i: number, j: number) {
    const rowLimit = myGrid.length - 1;
    const columnLimit = myGrid[0].length - 1;
    let xAmount = 0;
    for (let x = Math.max(0, i - 1); x <= Math.min(i + 1, rowLimit); x++) {
        for (let y = Math.max(0, j - 1); y <= Math.min(j + 1, columnLimit); y++) {
            if (x !== i || y !== j) {
                if ( myGrid[x][y] === "X") {
                    xAmount++;
                }
            }
        }
    }
    return xAmount;
}

    /*
    actual:
    [
        [ 'X', 'X', 'X', 'X', 'X', '0' ],
        [ '0', '0', '0', '0', '0', '0' ],
        [ '0', '0', '0', '0', '0', '0' ],
        [ '0', '0', '0', '0', '0', '0' ],
        [ '0', '0', '0', '0', '0', '0' ],
        [ '0', '0', '0', '0', '0', '0' ]
    ]

    expected:
    [
        ["X", "X", "X", "X", "X", "1"],
        [ "2", "3", "3", "3", "2", "1" ],
        [ "0", "0", "0", "0", "0", "0" ],
        [ "0", "0", "0", "0", "0", "0" ],
        [ "0", "0", "0", "0", "0", "0" ],
        [ "0", "0", "0", "0", "0", "0" ],
    ]
    */
