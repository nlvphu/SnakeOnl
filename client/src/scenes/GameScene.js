import Phaser from 'phaser'

export default class GameScene extends Phaser.Scene {
    constructor() {
      super({ key: 'GameScene' });
      this.socket = null;
      this.snakes = {};
    }
  
    preload() {
      this.load.image("food","public/images/apple.png");
      this.load.image("snakeRight","public/images/head_right.png");
      this.load.image("snakeLeft","public/images/head_left.png");
      this.load.image("snakeUp","public/images/head_up.png");
      this.load.image("snakeDown","public/images/head_down.png");
      this.load.image("bodyHorizontal","public/images/imbody_horizontal.png");
      this.load.image("bodyVertical","public/images/body_vertical.png");
      this.load.image("bodyRightUp","public/images/body_rightup.png");
      this.load.image("bodyRightDown","public/images/body_rightdown.png");
      this.load.image("bodyDownRight","public/images/body_downright.png");
      this.load.image("bodyUpRight","public/images/body_upright.png");
      this.load.image("tailRight","public/images/tail_right.png");
      this.load.image("tailLeft","public/images/tail_left.png");
      this.load.image("tailUp","public/images/tail_up.png");
      this.load.image("tailDown","public/images/tail_down.png");
      //this.load.image("background","assets/images/bg.png");
    }
  
    create() {
      // Connect to WebSocket server
      this.socket = new WebSocket('ws://localhost:8080/ws?id=phaserPlayer');
      
      // Handle connection open event
      this.socket.onopen = () => {
        console.log('WebSocket connection established');
      };

      // Handle incoming messages
      this.socket.onmessage = (event) => {
        const gameState = JSON.parse(event.data);
        this.updateGameState(gameState);
        console.log("Im right here");
      };

      // Handle connection close event
      this.socket.onclose = (event) => {
        console.log('WebSocket connection closed', event);
      };

      // Handle connection error event
      this.socket.onerror = (error) => {
        console.log('WebSocket error', error);
      };

  
      // Capture keyboard input for snake direction
      this.cursors = this.input.keyboard.createCursorKeys();
    }
  
    update() {
      if (this.cursors.left.isDown) {
        this.socket.send('LEFT');
      } else if (this.cursors.right.isDown) {
        this.socket.send('RIGHT');
      } else if (this.cursors.up.isDown) {
        this.socket.send('UP');
      } else if (this.cursors.down.isDown) {
        this.socket.send('DOWN');
      }
  
      // Render snakes based on game state
      for (const id in this.snakes) {
        const snake = this.snakes[id];
        if (!snake.graphics) {
          snake.graphics = this.add.graphics();
        }
        snake.graphics.clear();
        snake.graphics.fillStyle(0x00ff00, 1.0);
        snake.body.forEach(segment => {
          snake.graphics.fillRect(segment.X * 10, segment.Y * 10, 10, 10);
        });
      }
    }
  
    updateGameState(gameState) {
      this.snakes = gameState;
    }
  }
  