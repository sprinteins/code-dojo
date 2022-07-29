import { expect } from "chai";
import { getGrid } from ".";
import { getAmountOfNeighbourBombs } from ".";

describe("Scenarion: I want to receiver a minesweeper grid", () => {

    describe("WHEN I ask for a minesweeper GRID", () => {
        it("THEN I can see that i have a GRID", () => {
            const grid = getGrid(5, 5, 5);

            expect(grid.width).to.eq(5);
            expect(grid.height).to.eq(5);
        });

        describe("WHEN i ask for a 5 dimensional GRID", () => {
            it("THEN I can see that i receive a 5 deimensional GRID", () => {
                const width = 5;
                const height = 5;
                const grid = getGrid(width, height, 5);

                expect(grid.fields).to.have.length(width);
                expect(grid.fields[0]).to.have.length(height);
                expect(grid.fields[1]).to.have.length(height);
                expect(grid.fields[2]).to.have.length(height);
                expect(grid.fields[3]).to.have.length(height);
                expect(grid.fields[4]).to.have.length(height);
            });
        });

        describe("WHEN i ask for a 6 dimensional GRID", () => {
            it("THEN I can see that i receive a 6 deimensional GRID", () => {
                const width = 6;
                const height = 6;
                const grid = getGrid(width, height, 5);

                expect(grid.fields).to.have.length(width);
                expect(grid.fields[0]).to.have.length(height);
                expect(grid.fields[1]).to.have.length(height);
                expect(grid.fields[2]).to.have.length(height);
                expect(grid.fields[3]).to.have.length(height);
                expect(grid.fields[4]).to.have.length(height);
            });
        });

        describe("WHEN i ask for a 6 dimensional GRID", () => {
            it("THEN I can see that GRID", () => {
                const width = 6;
                const height = 6;
                const expectedBombAmount = 5;
                const grid = getGrid(width, height, expectedBombAmount);
                let actualBombAmount = 0;
                grid.fields.forEach((row) => {
                    actualBombAmount += row.filter((r) => r === "X").length;
                });

                console.log(grid);

                expect(actualBombAmount).to.eq(expectedBombAmount);
            });
        });

        describe("WHEN i ask for a 6 dimensional GRID", () => {
            it("THEN I can see that GRID", () => {
                const width = 6;
                const height = 6;
                const expectedBombAmount = 7;
                const grid = getGrid(width, height, expectedBombAmount);
                let actualBombAmount = 0;
                grid.fields.forEach((row) => {
                    actualBombAmount += row.filter((r) => r === "X").length;
                });

                expect(actualBombAmount).to.eq(expectedBombAmount);
            });
        });

        describe(`WHEN i ask for a 6 dimensional GRID
        AND I want to have 8 bombs on the field`, () => {
            it("THEN I can see that GRID", () => {
                const width = 6;
                const height = 6;
                const expectedBombAmount = 8;
                const grid = getGrid(width, height, expectedBombAmount);
                let actualBombAmount = 0;
                grid.fields.forEach((row) => {
                    actualBombAmount += row.filter((r) => r === "X").length;
                });

                expect(actualBombAmount).to.eq(expectedBombAmount);
            });
        });

        describe(`WHEN i ask for a 6 dimensional GRID
        AND I want to have 5 bombs on the field
        AND I can see that the fields around bombs are numbered`, () => {
            it("THEN I can see that GRID", () => {
                const width = 6;
                const height = 6;
                const bombAmount = 5;
                const grid = getGrid(width, height, bombAmount);
                const expectedField = [
                    ["X", "X", "X", "X", "X", "0"],
                    [ "0", "0", "0", "0", "0", "0" ],
                    [ "0", "0", "0", "0", "0", "0" ],
                    [ "0", "0", "0", "0", "0", "0" ],
                    [ "0", "0", "0", "0", "0", "0" ],
                    [ "0", "0", "0", "0", "0", "0" ],
                ];

                expect(grid.fields).to.deep.eq(expectedField);
            });
        });

        describe(`WHEN i ask for a 6 dimensional GRID
        AND I want to have 5 bombs on the field
        AND I can see that the fields around bombs are numbered`, () => {
            it("THEN I can see that GRID", () => {
                const width = 6;
                const height = 6;
                const bombAmount = 5;
                const grid = getGrid(width, height, bombAmount);
                const fields = grid.setFlags();
                const expectedField = [
                    ["X", "X", "X", "X", "X", "1"],
                    [ "2", "3", "3", "3", "2", "1" ],
                    [ "0", "0", "0", "0", "0", "0" ],
                    [ "0", "0", "0", "0", "0", "0" ],
                    [ "0", "0", "0", "0", "0", "0" ],
                    [ "0", "0", "0", "0", "0", "0" ],
                ];

                expect(fields).to.deep.eq(expectedField);
            });
        });

        describe(`WHEN `, () => {
            it("THEN I can see that GRID", () => {
                const fields = [
                    ["X", "X", "X"],
                    [ "X", "8", "X"],
                    [ "X", "X", "X"],
                ];
                const neighbouringBombs = getAmountOfNeighbourBombs(fields, 1, 1);

                expect(neighbouringBombs).to.deep.eq(8);
            });
        });
    });
});

/*
[
    [X,1,,,]
    [1,1,,,]
    [1,1,1,,]
    [2,X,2,,]
    [2,X,2,,]
]
*/
