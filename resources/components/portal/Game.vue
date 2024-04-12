<template>
  
    <canvas id="canvas1"></canvas>

</template>

<style>
  * {
    margin: 0;
    padding: 0;
    box-sizing: border-box;
  }

  #canvas1 {
    border: 5px solid black;
    position: absolute;
    top: 0;
    left: 0;
    background-color: blue;
    max-width: 100%;
    max-height: 100%;
  }
</style>

<script>
  define(function(){
    var component = {
      template: template,
      mounted: function() {

        function Game(canvas, context){
          this.canvas = canvas;
          this.context = context;
          this.width = this.canvas.width;
          this.height = this.canvas.height;
        }

        function Player(game){
          this.game = game;
          this.x = 20;
          this.y;
          this.spriteWidth = 200;
          this.spriteHeight = 200;
          this.width = 200;
          this.height = 200;
          this.speedY;
          this.flapSpeed;
        }
        Player.prototype.draw = function(){
          this.game.ctx.fillRect(this.x, this.y, this.width, this.height)
        }
        Player.prototype.update = function () {
          this.y += this.speedY;
          if (!this.isTouchingBottom()) {
            this.speedY += this.game.gravity;
          }
          // bottom boundary
          if (this.isTouchingBottom()) {
            this.y = this.game.height - this.height;
          }
        }
        Player.prototype.resize = function () {
          this.width = this.spriteWidth * this.game.ratio;
          this.height = this.spriteHeight * this.game.ratio;
          this.y = this.game.height * 0.5 - this.height * 0.5;
          this.speedY = -4 * this.game.ratio;
          this.flapSpeed = 7 * this.game.ratio;
        }
        Player.prototype.isTouchingTop = function () {
          return this.y <= -this.game.height * 0.1;
        }
        Player.prototype.isTouchingBottom = function () {
          return this.y >= this.game.height - this.height;
        }
        Player.prototype.flap = function () {
          if (!this.isTouchingTop()){
            this.speedY = -this.flapSpeed;
          };
        }
        

        function Obstacle(game, x) {
          this.game = game;
          this.spriteWidth = 110;
          this.spriteHeight = 110;
          this.scaledWidth = this.spriteHeight * this.game.ratio;
          this.scaledHeight = this.spriteHeight * this.game.ratio;
          this.x = x;
          this.y = Math.random() * this.game.height - this.scaledHeight;
          this.speedY = Math.random() < 0.5 ? -1 : 1;
        }
        Obstacle.prototype.update = function() {
          this.x -= this.game.speed * 1.5;
          this.y += this.speedY;
          if (this.y <= this.scaledHeight * -0.2 || this.y >= this.game.height - this.scaledHeight){
            this.speedY *= -1;
          }
        }
        Obstacle.prototype.draw = function() {
          this.game.ctx.fillRect(this.x, this.y, this.scaledWidth, this.scaledHeight);
        }
        Obstacle.prototype.resize = function() {
          this.scaledWidth = this.spriteWidth * this.game.ratio;
          this.scaledHeight = this.spriteHeight * this.game.ratio;
        }

        function Background(game) {

        }

          const canvas = document.getElementById('canvas1');
          const ctx = canvas.getContext('2d');
          canvas.width = 720;
          canvas.height = 720;
     
        var game = new Game(canvas, ctx);
      }
    };

    return component;
  })
</script>
