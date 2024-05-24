<template>
  <div id="canvas-container">
    <canvas id="canvas1"></canvas>
    <div id="highScore-text">0</div>

    <div class="assets">
      <img id="background" src='<% .Helpers.AssetPath "backgrounds/background trees with border.png" %>'>
      <img id="background2" src='<% .Helpers.AssetPath "backgrounds/background sky dark.png" %>'>
      <img id="PSprite" src='<% .Helpers.AssetPath "icons/UFO2.png" %>'>
      <img id="woodBoard" src='<% .Helpers.AssetPath "backgrounds/wood board.png" %>'>
      <img id="OSprite1" src='<% .Helpers.AssetPath "obstacles/obstacles.png" %>'>
      <img id="playerButton" src='<% .Helpers.AssetPath "backgrounds/buttons2.png" %>'>
      <button type="image">
        <img id="pause-btn" src='<% .Helpers.AssetPath "icons/PauseIcon.png" %>' alt="Pause">
      </button>
      <button type="image">
        <img id="score-board" src='<% .Helpers.AssetPath "backgrounds/ScoreBoard.png" %>' alt="Score Board">
      </button>
      <button type="image">
        <img id="pause-board" src='<% .Helpers.AssetPath "backgrounds/PauseBoard.png" %>' alt="Pause">
      </button>
      <div id="score-text">0</div>
      <div id="score-over">0</div>
    </div>
  </div>
</template>

<style>
body,
html {
  width: 100%;
  height: 100%;
  overflow: hidden;
}

#canvas-container {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
}

#canvas1 {
  width: 100%;
  height: 100%;
}

 #highScore-text {
   position: absolute;
   top: 2.5%;
   right: 3%;
   /* transform: translate(-50%, -50%); */
   font-size: 2.5vh;
   color: white;
   /* text-shadow: 1px 1px 2px black; */
 }

 #score-text {
  position: absolute;
  top: 20%;
  left: 50%;
  font-size: 8vh;
  color: white;
  transform: translate(-50%, -50%);
  text-shadow: 1.5px 1.5px 2px black;
  background: transparent;
  text-align: center;
  /* visibility: "visible"; */
 }

 #score-over {
   position: absolute;
   top: 38%;
   left: 50%;
   font-size: 7vh;
   color: black;
   transform: translate(-50%, -50%);
   background: transparent;
   text-align: center;
  visibility: hidden;
 }

 #score-board {
   position: absolute;
   top: 45%;
   left: 50%;
   transform: translate(-50%, -50%);
   height: 35vh;
   width: 35vh;
  visibility: hidden;
 }

  #pause-board {
    position: absolute;
    top: 45%;
    left: 50%;
    transform: translate(-50%, -50%);
    height: 35vh;
    width: 35vh;
    visibility: hidden;
  }

  #pause-btn {
    position: absolute;
    top: 2%;
    left: 3%;
    height: 5vh;
    width: 5vh;
    visibility: visible;
    /* transform: rotate(0deg); */

  }

</style>

<script>
define(function () {
  var component = {
    template: template,

    mounted: function () {
      // window.location.reload();

      function Game(canvas, context) {

        var self = this;

        this.canvas = canvas;
        this.ctx = context;
        this.width = this.canvas.width;
        this.height = this.canvas.height;
        this.baseHeight = 720;
        this.ratio = this.height / this.baseHeight;
        this.background = new Background(this);
        this.player = new Player(this);
        this.obstacles = [];
        this.obstacles2 = [];
        this.numberOfObstacles = 5;
        this.gravity;
        this.speed;
        this.minSpeed;
        this.maxSpeed;
        this.score = 0;
        this.highScore = 0;
        this.gameOver;
        this.timer = 0;
        this.pause = false;
        this.submitting = true; 
        this.difficulty = 0.8;  
        // this.turn = 1;

        this.woodImage = document.getElementById('woodBoard');
        this.buttonImage = document.getElementById('playerButton');
        this.highScoreDisplay = document.getElementById('highScore-text');
        this.scoreDisplay = document.getElementById('score-text');
        this.scoreBoard = document.getElementById('score-board');
        this.pauseBoard = document.getElementById('pause-board');
        this.boardScore = document.getElementById('score-over');
        this.pauseBtn = document.getElementById('pause-btn');

        var formQuery = { highScore: self.highScore };

        $flare.http.get('<% .Helpers.UrlForRoute "score.get" %>', formQuery)
          .then(function (response) {
            console.log(response.highScore);
            self.highScore = response.highScore;
            console.log(self.highScore);
          })
          .catch(function (error) {
            console.log(error);
          });
        
        this.resize(window.innerWidth, window.innerHeight);

        console.log("Game Started");
        console.log("Created Obstacles: " + this.numberOfObstacles);

        window.addEventListener("orientationchange", function () {
          window.location.reload();
        });

        this.scoreBoard.addEventListener("click", function () {
          window.location.reload();
        });

        this.pauseBtn.addEventListener("click", function () {
          self.pauseGame();
        });

        this.pauseBoard.addEventListener("click", function (){
          self.pauseGame();
        });

        window.addEventListener('resize', e => {
          this.resize(e.currentTarget.innerWidth, e.currentTarget.innerHeight);
        });

        this.canvas.addEventListener('mousedown', e => {
          if (!this.gameOver){
            this.player.flap(1);
          } else if (this.gameOver) {
            // window.location.reload();
          }
          console.log("mousedown");    
        });

        document.addEventListener('touchmove', function (e) {
          e.preventDefault();
        }, { passive: false });

        document.addEventListener('keydown', e => {
          if (e.key === "r" && this.gameOver) {
            console.log("r pressed");
            window.location.reload();
          }

          if (e.key === " " && !this.gameOver) {
            this.player.flap(1);
            console.log("spacebar pressed");
          }

          if (e.key === "p" && !this.gameOver) {
           this.pauseGame();
          }

          if (e.key === "e" && !this.gameOver) {
            if (this.player.energy >= this.player.maxEnergu*.75){
              if (this.player.charging) { 
                this.player.stopCharge(); 
              } else {
                this.player.startCharge();
              }       
            } 
          }
        });
      }
      Game.prototype.render = function (deltaTime) {

        // this.turn += 1;
        // this.pauseBtn.style.transform = "rotate(" + (this.turn % 360) + "deg)"

        if (!this.gameOver) {
          //create obstacle loop
          if (this.obstacles.length <= 3) this.addObstacles();

          if (!this.pause) {
            this.timer += deltaTime;
            
             };    

        };    

        if (this.gameOver) {
          //Menu during Game Over
          this.gameOverMenu();
          this.pauseBtn.style.visibility = "hidden";
        };

        if (!this.pause){
          this.background.update();
          this.player.update();
        }

        this.background.draw();
        this.player.draw();

        this.obstacles.forEach(obstacle => {
          if (!this.pause) obstacle.update();
          obstacle.draw();
        });
        this.drawStatusText();
        
      }
      Game.prototype.resize = function (width, height) {
        this.speed = 3 * this.ratio * this.difficulty;
        this.canvas.width = width;
        this.canvas.height = height;
        this.ctx.font = '20px Impact'
        this.ctx.textAlign = 'right';

        this.width = this.canvas.width;
        this.height = this.canvas.height;
        this.ratio = this.height / this.baseHeight;

        this.ctx.lineWidth = 3;
        this.ctx.strokeStyle = "white";

        this.gravity = 0.17 * this.ratio;
        
        this.minSpeed = this.speed;
        this.maxSpeed = this.speed * 1.5;

        this.background.resize();
        this.player.resize();

        this.createObstacles();
        this.obstacles.forEach(obstacle => {
          obstacle.resize();
        });

        this.gameOver = false;
        this.time = 0;
      }
      Game.prototype.createObstacles = function () {
        this.obstacles = [];
        const firstX = this.baseHeight * this.ratio;
        const obstacleSpacing = 800 * this.ratio;
        for (let i = 0; i < this.numberOfObstacles; i++) {
          this.obstacles.push(new Obstacle(this, firstX + i * obstacleSpacing));
        }
      }
      Game.prototype.addObstacles = function () {
        this.obstacles2 = [];
        const firstX = 4.4 * this.baseHeight * this.ratio;
        const obstacleSpacing = 800 * this.ratio;
        for (let i = 0; i < this.numberOfObstacles; i++) {
          this.obstacles.push(new Obstacle(this, firstX + i * obstacleSpacing));
        }
      }
      Game.prototype.checkCollision = function (a, b){
        const dx = a.collisionX - b.collisionX;
        const dy = a.collisionY - b.collisionY;
        const distance = Math.hypot(dx, dy);
        const sumOfRadii = a.collisionRadius + b.collisionRadius;
        return distance <= sumOfRadii;
      }
      Game.prototype.formatTimer = function () {
        return (this.timer * 0.001).toFixed(1);
      }
      Game.prototype.drawStatusText = function () {
        this.ctx.save();

        this.highScoreDisplay.innerText = "High Score: " + this.highScore;
        this.scoreDisplay.innerText = this.score
        if (this.gameOver) this.scoreDisplay.style.visibility = "hidden";

        // this.ctx.font = '25px Impact';
        // // this.ctx.fillText('Score: ' + this.score, this.width -10, 55);
        // this.ctx.fillText('High Score: ' + this.highScore, this.width - 10, 30);

        // this.ctx.textAlign = 'left';
        // this.ctx.fillText('Time: ' + this.formatTimer(this.timer), 10, 30);

        // if (!this.gameOver) {
        //   this.ctx.textAlign = 'center';
        //   this.ctx.font = '50px Impact';
        //   this.ctx.strokeStyle = 'black';
        //   this.ctx.fillStyle = 'white';
        //   this.ctx.strokeText(this.score, this.width * 0.5, this.height * 0.3)
        //   // this.ctx.font = '52px Impact';
        //   this.ctx.fillText(this.score, this.width * 0.5, this.height * 0.3)
        // }

        this.difficulty = 1 + this.timer / 60000;
        if (!this.player.charging) { this.speed = 3 * this.ratio * this.difficulty; }
      
  
        this.ctx.restore();
      }
      Game.prototype.gameOverMenu = function () {
        this.highScoreDisplay.style.visibility = "hidden";
        this.boardScore.style.visibility = "visible";
        this.boardScore.innerText = this.score;
        this.scoreBoard.style.visibility = "visible";
      }
      Game.prototype.pauseGame = function () {
        console.log("paused pressed");

        if (this.pause) {
          this.pause = false;
          this.pauseBtn.style.visibility = "visible";
          this.pauseBoard.style.visibility = "hidden";
          this.scoreDisplay.style.visibility = "visible";
          this.boardScore.style.visibility = "hidden"
          // requestAnimationFrame(animate);
        } else {
          this.pause = true;
          this.pauseBoard.style.visibility = "visible";
          this.pauseBtn.style.visibility = "hidden";
          this.scoreDisplay.style.visibility = "hidden";
          this.boardScore.innerText = this.score;
          this.boardScore.style.visibility = "visible"
        }
        console.log(this.pause);   
      }

      function Player(game) {
        this.game = game;
        this.x = 20;
        this.y;
        this.spriteWidth = 75;
        this.spriteHeight = 75;
        this.width = 50;
        this.height = 50;
        this.speedY;
        this.minSpeed;
        this.maxSpeed;
        this.flapSpeed;
        this.collisionX;
        this.collisionY;
        this.collisionRadius;
        this.collided = false;
        this.energy = 20;
        this.maxEnergu = this.energy * 3;
        this.minEnergy = 15;
        this.barSize;
        this.image = document.getElementById("PSprite");
        this.charging = false;
        this.invulnerable = false;
      }
      Player.prototype.draw = function () {
        // this.game.ctx.strokeRect(this.x, this.y, this.width, this.height);
        this.game.ctx.drawImage(this.image, this.x, this.y, this.width, this.height)
        this.game.ctx.beginPath();
        this.game.ctx.arc(this.collisionX, this.collisionY, this.collisionRadius, 0, Math.PI * 2);
        if (this.charging) this.game.ctx.stroke();
      }
      Player.prototype.update = function () {
        this.handleEnergy();
        this.y += this.speedY;
        this.collisionY = this.y + this.height * 0.45;
        if (!this.isTouchingBottom()) {
          this.speedY += this.game.gravity;
        }
        if (this.isTouchingTop()) {
          this.speedY = this.flapSpeed*0.5;
        }
        // bottom boundary
        if (this.isTouchingBottom() && !this.game.gameOver) {          

         this.flap(0.6);

        }
       
        if(this.game.gameOver){
          if (this.x >= -300){
          this.x -= 1;
          }
        }
      }
      Player.prototype.resize = function () {
        this.width = this.spriteWidth * this.game.ratio;
        this.height = this.spriteHeight * this.game.ratio;
        this.y = this.game.height * 0.5 - this.height * 0.5;
        this.speedY = -4 * this.game.ratio;
        this.flapSpeed = 7 * this.game.ratio;
        this.collisionRadius = this.width * 0.5 * 0.7;
        this.collisionX = this.x + this.width * 0.5;
        this.collided = false;
        this.barSize = 5 * this.game.ratio;
        // this.collisionY = this.y + this.height * 0.5;
      }
      Player.prototype.isTouchingBottom = function () {
        return this.y >= this.game.height - this.game.height*0.32;
      }
      Player.prototype.isTouchingTop = function () {
        return this.y <= this.game.height * 0.1;
      }
      Player.prototype.flap = function (x) {
        // this.stopCharge();
        if (!this.isTouchingTop()) {
          this.speedY = -this.flapSpeed * x;
        };
      }
      Player.prototype.handleEnergy = function () {
        if (this.energy < this.maxEnergu && !this.charging && !this.game.gameOver){
          this.energy += 0.1;
        }
        if (this.energy <= 0){
          this.energy = 0;
          this.stopCharge();
        }
        if (this.charging) {
          this.energy -= 0.2;
        }
      }
      Player.prototype.startCharge = function () {
   
        this.charging = true;
        this.game.speed = this.game.maxSpeed;
      }
      Player.prototype.stopCharge = function () { 
        this.charging = false;
        this.game.speed = this.game.minSpeed;
      }
      Player.prototype.isInvulnerable = function (){
        if (this.charging){
          return true
        }
      }
      // Player.prototype.shield = function () {
      //   //this.game.ctx.strokeRect(this.x, this.y, this.width, this.height);
      //   // this.game.ctx.drawImage(this.image, this.x, this.y, this.width, this.height)
      //     if (this.invulnerable){
      //     this.game.ctx.beginPath();
      //     this.game.ctx.arc(this.collisionX, this.collisionY, this.collisionRadius, 0, Math.PI * 2);
      //     this.game.ctx.stroke();
      //   }
      // }

      function Background(game) {
        this.game = game;
        this.image = document.getElementById('background');
        this.image2 = document.getElementById('background2');
        this.width = 2400;
        this.height = this.game.baseHeight;
        this.scaledWidth;
        this.scaledHeight;
        this.y;
        this.x;
        this.x2;
      }
      Background.prototype.update = function () {
        if(this.game.gameOver){
          this.game.speed = 0.2;
        }

        this.x -= this.game.speed * 1;
        if (this.x <= -this.scaledWidth) this.x = 0;
        this.x2 -= this.game.speed * 0.5;
        if (this.x2 <= -this.scaledWidth) this.x2 = 0;
      }
      Background.prototype.draw = function () {
        this.game.ctx.drawImage(this.image2, this.x2, this.y, this.scaledWidth, this.scaledHeight);
        this.game.ctx.drawImage(this.image2, this.x2 + this.scaledWidth - 1, this.y, this.scaledWidth, this.scaledHeight);
        if (this.game.canvas.width >= this.scaledWidth) {
          this.game.ctx.drawImage(this.image2, this.x2 + this.scaledWidth - 1, this.y, this.scaledWidth, this.scaledHeight);
        }
        this.game.ctx.drawImage(this.image, this.x, this.y, this.scaledWidth, this.scaledHeight);
        this.game.ctx.drawImage(this.image, this.x + this.scaledWidth - 1, this.y, this.scaledWidth, this.scaledHeight);
        if (this.game.canvas.width >= this.scaledWidth) {
          this.game.ctx.drawImage(this.image, this.x + this.scaledWidth - 1, this.y, this.scaledWidth, this.scaledHeight);
        }
      }
      Background.prototype.resize = function () {
        this.scaledWidth = this.width * this.game.ratio;
        this.scaledHeight = this.height * this.game.ratio;
        this.y = 0;
        this.x = 0;
        this.x2 = 0;
      }

      function Obstacle(game, x) {
        this.game = game;
        this.spriteWidth = 200;
        this.spriteHeight = 200;
        this.width = 120;
        this.height = 120;
        this.scaledWidth = this.width * this.game.ratio;
        this.scaledHeight = this.height * this.game.ratio;
        this.x = x;
        this.y = this.game.height * 0.2 + this.game.height * (0.4 * Math.random());
        this.collisionX;
        this.collisionY;
        this.speedY = Math.random() < 0.5 ? -1 : 1;
        this.markedFordeletion = false;
        this.image = document.getElementById("OSprite1");
        this.randomPic = Math.floor(4 * Math.random());
        this.collisionRadius = this.scaledWidth * 0.5 * 0.8;
        this.collided = false;
    
      }
      Obstacle.prototype.update = function () {
        
        this.y += this.speedY;
        this.x -= this.game.speed * 1.5;

        if (!this.game.gameOver){                   
          this.collisionY = this.y + this.scaledHeight * 0.5;
          if (this.y <= this.game.height * 0.1 || this.y >= this.game.height * 0.65) {
            this.speedY *= -1;
          }
        } else { this.y += this.speedY * 2}

        this.collisionX = this.x + this.scaledWidth * 0.5;

        if (this.isOffScreen() || this.collided) {
          this.markedFordeletion = true;
          this.game.obstacles = this.game.obstacles.filter(obstacle => !obstacle.markedFordeletion);
          console.log("obstacles left: " + this.game.obstacles.length);
          if (!this.game.gameOver) this.game.score++;
        }
        if (this.game.obstacles.length <= 0) {
          this.game.gameOver = true;
        }
        //collision
        if (!this.game.player.isInvulnerable()){
        if (this.game.checkCollision(this, this.game.player)){
        
          this.game.gameOver = true;
          this.game.player.collided = true;
          this.collided = true;
          
        }
        }
      }
      Obstacle.prototype.draw = function () {
        // this.game.ctx.fillRect(this.x, this.y, this.scaledWidth, this.scaledHeight);
        this.game.ctx.drawImage(this.image, 0, this.randomPic * this.spriteHeight, this.spriteHeight, this.spriteHeight, this.x, this.y, this.scaledWidth, this.scaledHeight);
        this.game.ctx.beginPath();
        this.game.ctx.arc(this.collisionX, this.collisionY, this.collisionRadius, 0, Math.PI * 2);
        //this.game.ctx.stroke();
      }
      Obstacle.prototype.resize = function () {
        this.scaledWidth = this.width * this.game.ratio;
        this.scaledHeight = this.height * this.game.ratio;
      }
      Obstacle.prototype.isOffScreen = function () {
        return this.x < -this.scaledWidth;
      }

      const canvas = document.getElementById('canvas1');
      const ctx = canvas.getContext('2d');
      canvas.width = 300;
      canvas.height = 300;
      const game = new Game(canvas, ctx);
      game.render(0);

      let lastTime = 0;
      function animate(timeStamp) {
        const deltaTime = timeStamp - lastTime;
        lastTime = timeStamp;
        ctx.clearRect(0, 0, canvas.width, canvas.height);
        game.render(deltaTime);
        // if (!game.pause) requestAnimationFrame(animate);
        requestAnimationFrame(animate);
        // requestAnimationFrame(animate);
      }
      requestAnimationFrame(animate);
    }

  };

  return component;
})
</script>
