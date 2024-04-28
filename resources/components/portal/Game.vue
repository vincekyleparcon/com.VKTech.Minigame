<template>
  <div>
    <canvas id="canvas1"></canvas>
    <div class="assets">
      <img id="background" src='<% .Helpers.AssetPath "backgrounds/background trees.png" %>'>
      <img id="background2" src='<% .Helpers.AssetPath "backgrounds/background sky.png" %>'>
      <img id="PSprite" src='<% .Helpers.AssetPath "icons/UFO2.png" %>'>
      <img id="OSprite1" src='<% .Helpers.AssetPath "obstacles/obstacles.png" %>'>
    </div>
  </div>
</template>

<style>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
  font-family: Impact, Haettenschweiler, 'Arial Narrow Bold', sans-serif;
}

#canvas1 {
  position: fixed;
  top: 0;
  left: 0;
  max-width: 100%;
  max-height: 100%;
  overflow: hidden;
}

</style>

<script>
define(function () {
  var component = {
    template: template,
    mounted: function () {

      function Game(canvas, context) {
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
        this.numberOfObstacles = 8;
        this.gravity;
        this.speed;
        this.minSpeed;
        this.maxSpeed;
        this.score = 0;
        this.highScore = 0;

        // var formData = { 
        //   name: "Kyle",
        //   score: this.score,  
        // };
        
        var formData = { highScore: this.highScorescore };
          $flare.http.get('<% .Helpers.UrlForRoute "score.get" %>', formData)
            .then(function (response) {
              console.log(response);
            })
            .catch(function (error) {
              console.log(error);
            });

        

        this.gameOver;
        this.timer = 0;

        this.resize(window.innerWidth, window.innerHeight);

        console.log("Game Started");
        console.log("Created Obstacles: " + this.numberOfObstacles);

        window.addEventListener('resize', e => {
          this.resize(e.currentTarget.innerWidth, e.currentTarget.innerHeight);
        });

        this.canvas.addEventListener('mousedown', e => {
          if (!this.gameOver){
            this.player.flap(1);
          }
          console.log("mousedown");    
        });

        document.addEventListener('keydown', e => {
          if (e.key === " " && !this.gameOver) {
            this.player.flap(1);
            console.log("spacebar pressed");
          }
        });
      }
      Game.prototype.render = function (deltaTime) {

       
        // console.log("DELTA: " + deltaTime);
        // this.timer += deltaTime;
        // console.log("TIMER: " + this.timer);
        if(this.obstacles.length <= 5) this.addObstacles();

        if (!this.gameOver) this.timer += deltaTime;

        this.background.update();
        this.background.draw();
        this.player.update();
        this.player.draw();

        this.drawStatusText();

        this.obstacles.forEach(obstacle => {
          obstacle.update();
          obstacle.draw();
        });
      }
      Game.prototype.resize = function (width, height) {
        this.canvas.width = width;
        this.canvas.height = height;
        this.ctx.fillStyle = 'orangered';
        this.ctx.font = '20px Impact'
        this.ctx.textAlign = 'right';

        this.width = this.canvas.width;
        this.height = this.canvas.height;
        this.ratio = this.height / this.baseHeight;

        this.ctx.lineWidth = 3;
        this.ctx.strokeStyle = "white";

        this.gravity = 0.17 * this.ratio;
        this.speed = 3 * this.ratio;
        this.background.resize();
        this.player.resize();

        this.createObstacles();
        this.obstacles.forEach(obstacle => {
          obstacle.resize();
        });

        this.gameOver = false;
        this.time = 0;
        this.isSubmitted = false;
        this.submitting = true;
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
        const firstX = 6.4 * this.baseHeight * this.ratio;
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

        this.ctx.font = '25px Impact';
        this.ctx.fillText('Score: ' + this.score, this.width -10, 55);
        this.ctx.fillText('High Score: ' + this.highScore, this.width - 10, 30);

        this.ctx.textAlign = 'left';
        this.ctx.fillText('Time: ' + this.formatTimer(this.timer), 10, 30);
        if (this.gameOver) {

          if (this.player.collided){
            this.message1 = "Game Over";
            this.message2 = "Score: " + this.score;

            var formData = { 
              score: this.score, 
              highScore: 999,
            };

            if (this.submitting){

              $flare.http.post('<% .Helpers.UrlForRoute "score.save" %>', formData)
                .then(function (response) {
                  console.log(response);
                })
                .catch(function (error) {
                  console.log(error);
                });
                this.submitting = false;

            }

          } else {
            this.message1 = "Win";
            this.message2 = "Score: " + this.score;
          }

          this.ctx.textAlign = 'center';
          this.ctx.font = '50px Impact';
          this.ctx.fillStyle = 'firebrick';
          this.ctx.fillText(this.message1, this.width * 0.5, this.height * 0.5 - 50)

          this.ctx.textAlign = 'center';
          this.ctx.font = '20px Impact';
          this.ctx.fillStyle = 'firebrick';
          this.ctx.fillText(this.message2, this.width * 0.5, this.height * 0.5 - 20)

          this.ctx.textAlign = 'center';
          this.ctx.font = '20px Impact';
          this.ctx.fillStyle = 'firebrick';
          this.ctx.fillText("Press R to try again", this.width * 0.5, this.height * 0.5 + 10)
        };
        
        
        for (let i = 0; i < this.player.energy; i++){   
          if (this.player.energy >= this.player.maxEnergu) this.ctx.fillStyle = 'yellow';    
          this.ctx.fillRect(10 + i*4, 40, 4, 15);
        }

        this.ctx.restore();
      }

      function Player(game) {
        this.game = game;
        this.x = 20;
        this.y;
        this.spriteWidth = 200;
        this.spriteHeight = 200;
        this.width = 200;
        this.height = 200;
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
        this.image = document.getElementById("PSprite");
      }
      Player.prototype.draw = function () {
        //this.game.ctx.strokeRect(this.x, this.y, this.width, this.height);
        this.game.ctx.drawImage(this.image, this.x, this.y, this.width, this.height)
        this.game.ctx.beginPath();
        this.game.ctx.arc(this.collisionX, this.collisionY, this.collisionRadius, 0, Math.PI * 2);
       // this.game.ctx.stroke();
      }
      Player.prototype.update = function () {
        this.handleEnergy();
        this.y += this.speedY;
        this.collisionY = this.y + this.height * 0.45;
        if (!this.isTouchingBottom()) {
          this.speedY += this.game.gravity;
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
        // this.collisionY = this.y + this.height * 0.5;
      }
      Player.prototype.isTouchingBottom = function () {
        return this.y >= this.game.height - this.height * 1.05;
      }
      Player.prototype.isTouchingTop = function () {
        return this.y <= -this.game.height * 0.1;
      }
      Player.prototype.flap = function (x) {
        if (!this.isTouchingTop()) {
          this.speedY = -this.flapSpeed * x;
        };
      }
      Player.prototype.handleEnergy = function () {
        if (this.energy < this.maxEnergu){
          this.energy += 0.75;
        }
      }
      Player.prototype.startCharge = function () {

      }
      Player.prototype.stopCharge = function () { 

      }

      function Background(game) {
        this.game = game;
        this.image = document.getElementById('background');
        this.image2 = document.getElementById('background2');
        this.width = 2400;
        this.height = this.game.baseHeight;
        this.scaledWidth;
        this.scaledHeight;
        this.x;
        this.x2;
      }
      Background.prototype.update = function () {
        if(this.game.gameOver){
          this.game.speed = 0.5;
        }

        this.x -= this.game.speed * 2;
        if (this.x <= -this.scaledWidth) this.x = 0;
        this.x2 -= this.game.speed * 1;
        if (this.x2 <= -this.scaledWidth) this.x2 = 0;
      }
      Background.prototype.draw = function () {
        this.game.ctx.drawImage(this.image2, this.x2, 0, this.scaledWidth, this.scaledHeight);
        this.game.ctx.drawImage(this.image2, this.x2 + this.scaledWidth - 1, 0, this.scaledWidth, this.scaledHeight);
        if (this.game.canvas.width >= this.scaledWidth) {
          this.game.ctx.drawImage(this.image2, this.x2 + this.scaledWidth * 2 - 2, 0, this.scaledWidth, this.scaledHeight);
        }
        this.game.ctx.drawImage(this.image, this.x, 0, this.scaledWidth, this.scaledHeight);
        this.game.ctx.drawImage(this.image, this.x + this.scaledWidth - 1, 0, this.scaledWidth, this.scaledHeight);
        if (this.game.canvas.width >= this.scaledWidth) {
          this.game.ctx.drawImage(this.image, this.x + this.scaledWidth * 2 - 2, 0, this.scaledWidth, this.scaledHeight);
        }
      }
      Background.prototype.resize = function () {
        this.scaledWidth = this.width * this.game.ratio;
        this.scaledHeight = this.height * this.game.ratio;
        this.x = 0;
        this.x2 = 0;
      }

      function Obstacle(game, x) {
        this.game = game;
        this.spriteWidth = 200;
        this.spriteHeight = 200;
        this.scaledWidth = this.spriteWidth * this.game.ratio;
        this.scaledHeight = this.spriteHeight * this.game.ratio;
        this.x = x;
        this.y = this.game.height * (0.5 * Math.random());
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
        this.x -= this.game.speed * 2.5;

        if (!this.game.gameOver){                   
          this.collisionY = this.y + this.scaledHeight * 0.5;
          if (this.y <= this.scaledHeight * -0.2 || this.y >= this.game.height - this.scaledHeight) {
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
        if (this.game.checkCollision(this, this.game.player)){
          this.game.gameOver = true;
          this.game.player.collided = true;
          this.collided = true;
        }
      }
      Obstacle.prototype.draw = function () {
        // this.game.ctx.fillRect(this.x, this.y, this.scaledWidth, this.scaledHeight);
        this.game.ctx.drawImage(this.image, 0, this.randomPic * this.spriteHeight, this.spriteWidth, this.spriteHeight, this.x, this.y, this.scaledWidth, this.scaledHeight);
        this.game.ctx.beginPath();
        this.game.ctx.arc(this.collisionX, this.collisionY, this.collisionRadius, 0, Math.PI * 2);
        //this.game.ctx.stroke();
      }
      Obstacle.prototype.resize = function () {
        this.scaledWidth = this.spriteWidth * this.game.ratio;
        this.scaledHeight = this.spriteHeight * this.game.ratio;
      }
      Obstacle.prototype.isOffScreen = function () {
        return this.x < -this.scaledWidth;
      }

      const canvas = document.getElementById('canvas1');
      const ctx = canvas.getContext('2d');
      canvas.width = 720;
      canvas.height = 720;
      const game = new Game(canvas, ctx);
      game.render(0);

      let lastTime = 0;
      function animate(timeStamp) {
        const deltaTime = timeStamp - lastTime;
        lastTime = timeStamp;
        ctx.clearRect(0, 0, canvas.width, canvas.height);
        game.render(deltaTime);
        //if (!game.gameOver) requestAnimationFrame(animate);
        requestAnimationFrame(animate);
      }
      requestAnimationFrame(animate);
    }

  };

  return component;
})
</script>
