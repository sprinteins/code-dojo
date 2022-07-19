
interface Grid {
    width: number;
    height: number;
    fields: string[][];
}

export const getGrid = (width: number, height: number, bombAmount: number): Grid => {
    
    let bombsPlanted = 0;
    var fields: string[][] = new Array(width).fill("0").map(() => {

        const row = new Array(height).fill("0")

        row.forEach((_, index) => {
           if (bombsPlanted < bombAmount) {
               row[index] = "X";
               bombsPlanted += 1;
           }
       })

        return row;
    });

    return {
        width,
        height,
        fields
    }
}
