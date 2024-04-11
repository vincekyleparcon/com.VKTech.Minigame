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
