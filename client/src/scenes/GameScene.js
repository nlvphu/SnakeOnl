import Phaser from 'phaser'

export default class GameScene extends Phaser.Scene {
    constructor() {
      super({ key: 'GameScene' });
      this.socket = null;
      this.snakes = {};
    }
  
    preload() {
      this.load.image("food","assets/apple.png");
      this.load.image("snakeRight","assets/head_right.png");
      this.load.image("snakeLeft","assets/head_left.png");
      this.load.image("snakeUp","assets/head_up.png");
      this.load.image("snakeDown","assets/head_down.png");
      this.load.image("bodyHorizontal","assets/body_horizontal.png");
      this.load.image("bodyVertical","assets/body_vertical.png");
      this.load.image("bodyTopLeft","assets/body_topleft.png");
      this.load.image("bodyTopRight","assets/body_topright.png");
      this.load.image("bodyBottomLeft","assets/body_bottomleft.png");
      this.load.image("bodyBottomRight","assets/body_bottomright.png");
      this.load.image("tailRight","assets/tail_right.png");
      this.load.image("tailLeft","assets/tail_left.png");
      this.load.image("tailUp","assets/tail_up.png");
      this.load.image("tailDown","assets/tail_down.png");
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
        console.log("On ONNNNNNNN message");
        const gameState = JSON.parse(event.data);
        this.updateGameState(gameState);
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
        snake.body.forEach(segment => {
          snake.graphics.fillStyle(0x00ff00, 1.0);
          snake.graphics.fillRect(segment.X * 10, segment.Y * 10, 10, 10);
        });
      }
    }
  
    updateGameState(gameState) {
      this.snakes = gameState;
    }
  }
  