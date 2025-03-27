/**
 * Represents a game's state.
 * @param {Player} player1
 * @param {Player} player2
 * @param {Target[]} targets
 * @param {Projectile[]} projectile
 */
export class GameState {
    constructor(player1, player2, targets, projectiles) {
        this.player1 = player1;
        this.player2 = player2;
        this.targets = targets;
    }
}

/**
 * Represents a player's state.
 * @param {number} x
 * @param {number} y
 * @param {number} health
 */
export class Player {
    constructor(x, y, health) {
        this.x = x
        this.y = y
        this.health = health;
    }
}


/**
 * Represents a target's state.
 * @param {number} x
 * @param {number} y
 * @param {number} velocity
 * @param {boolean} isEnabled
 * @param {number} diameter
 */
export class Target {
    constructor(x, y, velocity, isEnabled, diameter) {
        this.x = x
        this.y = y
        this.diameter = diameter
        this.velocity = velocity;
        this.isEnabled = isEnabled;
    }
}


/**
 * Represents a target's state.
 * @param {number} x
 * @param {number} y
 * @param {number} velocity
 * @param {boolean} isEnabled
 * @param {number} diameter
 */
export class Projectile {
    constructor(x, y, velocity, isEnabled, diameter) {
        this.x = x
        this.y = y
        this.diameter = diameter
        this.velocity = velocity;
        this.isEnabled = isEnabled;
    }
}
