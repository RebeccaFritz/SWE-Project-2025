/**
 * Represents a game's state.
 * @param {Player} player1
 * @param {Player} player1
 * @param {Target[]} player1
 */
export class GameState {
    constructor(player1, player2, targets) {
        this.player1 = player1;
        this.player2 = player2;
        this.targets = targets;
    }
}


/**
 * Represents a player's state.
 * @param {[number, number]} position
 * @param {number} health
 */
export class Player {
    constructor(position, health) {
        this.position = position;
        this.health = health;
    }
}


/**
 * Represents a target's state.
 * @param {[number, number]} position
 * @param {number} velocity
 * @param {boolean} isEnabled
 */
class Target {
    constructor(position, velocity, isEnabled) {
        this.position = position;
        this.velocity = velocity;
        this.isEnabled = isEnabled;
    }
}