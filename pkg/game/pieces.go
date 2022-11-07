package game

var out2 string = `
import { Square } from "./square"

// Piece movement related constants & interfaces
interface DirectionVec {
    file: number;
    rank: number;
}

const ORTHOGONAL: DirectionVec[] = [
    { file: 0, rank: -1 }, // North
    { file: 0, rank: 1 }, // South
    { file: -1, rank: 0 }, // West
    { file: 1, rank: 0 }, // East
];

const DIAGONAL: DirectionVec[] = [
    { file: -1, rank: -1 }, // Topleft
    { file: -1, rank: 1 }, // Topright
    { file: 1, rank: -1 }, // Bottomleft
    { file: 1, rank: 1 }, // Bottomright
];

const COMBINED: DirectionVec[] = [
    ...ORTHOGONAL,
    ...DIAGONAL
];

const KNIGHT: DirectionVec[] = [
    { file: -1, rank: -2 }, // TL
    { file: -2, rank: -1 }, // TL2
    { file: -2, rank: 1 }, // BL
    { file: -1, rank: 2 }, // BL2
    { file: 1, rank: -2 }, // TR
    { file: 2, rank: -1 }, // TR2
    { file: 2, rank: 1 }, // BR
    { file: 1, rank: 2 }, // BR2
];

const PIECE_ABBREVIATIONS: string[] = ["B", "R", "Q", "K", "N"];


/** Abstract base class for all piece types. */
export abstract class Piece {
    /* TODO in theory these should be static but if I need to access them in a 
    subclass I need to access them by class name instead of 'this.' */
    directions: DirectionVec[];
    abbreviation: string;

    square: Square;
    // pointer to board object used to get info about occupied squares
    readonly board: any;

    constructor(square: Square, board: any) {
        this.square = square;
        this.board = board;
    }

    abstract getMoves(): Square[]
}

abstract class SlidingPiece extends Piece {
    /** Get an array of squares to which the sliding piece instance can move to. */
    getMoves(): Square[] {
        const moves: Square[] = [];

        let new_square: Square;
        for (const direction of this.directions) {
            new_square = this.square;

            while (true) {
                const new_file = new_square.file + direction.file;
                const new_rank = new_square.rank + direction.rank;

                try {
                    new_square = new Square(new_file, new_rank);
                }
                catch (e) {
                    // we are out of bounds -> this square doesn't exist
                    // move to next direction
                    break;
                }

                // if square exists but its occupied -> move to next direction
                if (this.board.isOccupied(new_square)) {
                    break;
                }

                moves.push(new_square);
            }
        }

        return moves;
    }
}

abstract class NonSlidingPiece extends Piece {
    /** Get an array of squares to which the non-sliding piece instance can move to. */
    getMoves(): Square[] {
        const moves: Square[] = [];

        for (const direction of this.directions) {
            const new_file = this.square.file + direction.file;
            const new_rank = this.square.rank + direction.rank;

            let new_square: Square;
            try {
                new_square = new Square(new_file, new_rank);
            }
            catch (e) {
                // we are out of bounds -> this square doesn't exist
                // move to next direction
                continue;
            }

            // if square exists but its occupied -> move to next direction
            if (this.board.isOccupied(new_square)) {
                continue;
            }

            moves.push(new_square);
        }

        return moves;
    }
}

export class Bishop extends SlidingPiece {
    directions = DIAGONAL;
    abbreviation = "B";
}

export class Rook extends SlidingPiece {
    directions = ORTHOGONAL;
    abbreviation = "R";
}

export class Queen extends SlidingPiece {
    directions = COMBINED;
    abbreviation = "Q";
}

export class King extends NonSlidingPiece {
    directions = COMBINED;
    abbreviation = "K";
}

export class Knight extends NonSlidingPiece {
    directions = KNIGHT;
    abbreviation = "N";
} 
`
