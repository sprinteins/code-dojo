import { expect } from "chai";
import { getGrid } from ".";


describe('Scenarion: I want to receiver a minesweeper grid', () => {

    describe('WHEN I ask for a minesweeper GRID', () => {
        it('THEN I can see that i have a GRID', () => {
            const grid = getGrid(5, 5, 5);

            expect(grid.width).to.eq(5);
            expect(grid.height).to.eq(5);
        })

        describe('WHEN i ask for a 5 dimensional GRID', () => {
            it('THEN I can see that i receive a 5 deimensional GRID', () => {
                const width = 5;
                const height = 5;
                const grid = getGrid(width, height, 5);

                expect(grid.fields).to.have.length(width);
                expect(grid.fields[0]).to.have.length(height);
                expect(grid.fields[1]).to.have.length(height);
                expect(grid.fields[2]).to.have.length(height);
                expect(grid.fields[3]).to.have.length(height);
                expect(grid.fields[4]).to.have.length(height);
            })
        })

        describe('WHEN i ask for a 6 dimensional GRID', () => {
            it('THEN I can see that i receive a 6 deimensional GRID', () => {
                const width = 6;
                const height = 6;
                const grid = getGrid(width, height, 5);

                expect(grid.fields).to.have.length(width);
                expect(grid.fields[0]).to.have.length(height);
                expect(grid.fields[1]).to.have.length(height);
                expect(grid.fields[2]).to.have.length(height);
                expect(grid.fields[3]).to.have.length(height);
                expect(grid.fields[4]).to.have.length(height);
            })
        })

        describe('WHEN i ask for a 6 dimensional GRID', () => {
            it('THEN I can see that GRID', () => {
                const width = 6;
                const height = 6;
                const expectedBombAmount = 5;
                const grid = getGrid(width, height, expectedBombAmount);
                let actualBombAmount = 0;
                grid.fields.forEach((row) => {
                    actualBombAmount += row.filter((r) => r === "X").length;
                })
 
                console.log(grid);

                expect(actualBombAmount).to.eq(expectedBombAmount);
            })
        })

        describe('WHEN i ask for a 6 dimensional GRID', () => {
            it('THEN I can see that GRID', () => {
                const width = 6;
                const height = 6;
                const expectedBombAmount = 7;
                const grid = getGrid(width, height, expectedBombAmount);
                let actualBombAmount = 0;
                grid.fields.forEach((row) => {
                    actualBombAmount += row.filter((r) => r === "X").length;
                })
 
                expect(actualBombAmount).to.eq(expectedBombAmount);
            })
        })

        describe('WHEN i ask for a 6 dimensional GRID', () => {
            it('THEN I can see that GRID', () => {
                const width = 6;
                const height = 6;
                const expectedBombAmount = 38;
                const grid = getGrid(width, height, expectedBombAmount);
                let actualBombAmount = 0;
                grid.fields.forEach((row) => {
                    actualBombAmount += row.filter((r) => r === "X").length;
                })
 
                expect(actualBombAmount).to.eq(expectedBombAmount);
            })
        })
    })
})

/*
[
    [X,1,,,]
    [1,1,,,]
    [1,1,1,,]
    [2,X,2,,]
    [2,X,2,,]
]
*/