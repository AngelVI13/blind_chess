package game

var out string = `
import { Piece } from "./pieces"
import { Square } from "./square"


export class Board {
    pieces: Piece[];

    constructor() {
        this.pieces = [];
    }

    /** Resets board (i.e. removes all pieces from the board). */
    reset(): void {
        this.pieces = [];
    }

    /** Add a piece to the board */
    addPiece(piece: Piece): void {
        this.pieces.push(piece);
    }

    /** Checks if a given square is already occupied by a piece or not. */
    isOccupied(square: Square): boolean {
        const occupiedSquares = this.pieces.map(piece => piece.square);
        return square.isContained(occupiedSquares);
    }

    /** Return an array of squares to which only 1 piece can go. */
    getSingularSquares(): Square[] {
        // convert move arrays for each piece into a set
        const allMoves = this.pieces.map(piece => new Set(piece.getMoves()));
        
        // Calculate the symmetric difference between all moves.
        // This results in a set of squares which cannot be reached by more than one piece.
        const singularSquaresSet: Set<Square> = allMoves.reduce(symmetricSquareDifference);

        return Array.from(singularSquaresSet);
    }

    /** Get piece object that can reach the given square. */
    getPieceThatReachesSquare(square: Square): Piece {
        for (const piece of this.pieces) {
            if (square.isContained(piece.getMoves())) {
                return piece;
            }
        }
        // TODO: should this throw error or return null ?
        return null;
    }
}

/** 
 * Calculates symmetric difference between 2 sets of squares. 
 * Returns the squares which are not overlapping between the two input sets. 
 */
const symmetricSquareDifference = (accumulator: Set<Square>, currentValue: Set<Square>, currentIndex, array) => {
    // this implementation is taken directly from
    // https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Set
    let _difference = accumulator;

    // conversion to array is needed in order to allow iteration over elements
    for (let elem of Array.from(currentValue)) {
        // convert to array for iteration and containment checks
        const _differenceArray = Array.from(_difference);

        // need to use specific check for presence because we are dealing with square objects
        if (elem.isContained(_differenceArray)) {

            // iterate over all objects in accumulator array and delete
            // the one that is contained in the currentValue
            // NOTE: need to do a custom comparison by index
            //       because delete() will not find the matching
            //       object to delete as we are dealing with non-basic types
            for (const sq of _differenceArray) {
                if (sq.index() === elem.index()) {
                    _difference.delete(sq);
                }
            }

        } else {
            _difference.add(elem);
        }
    }
    return _difference;
}
`
