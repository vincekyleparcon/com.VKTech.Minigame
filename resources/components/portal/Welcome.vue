<template>


<!-- 
  <div class="mx-auto p-2">
    <form @click.prevent="submit">
      <label>Amount:</label>
      <input type="text" v-model="amount">
      <button type="submit">Submit</button>
    </form>

    <h1>Welcome {{ flareView.data.name }}</h1>

    <ul class="list-group list-group-horizontal mb-2" style="width: 90vw; max-width: 45rem;">
      <input type="text" class="form-control mr-2" placeholder="Username" aria-label="Username"
        aria-describedby="basic-addon1">
      <input type="text" class="form-control" placeholder="Username" aria-label="Username"
        aria-describedby="basic-addon1">
    </ul>

    <button type="button" class="btn btn-success mb-3">Sign In</button>

    <div class="rounded-lg p-4 mx-auto"
      style="width: 90vw; max-width: 45rem; height: 30rem; background-color: antiquewhite;">
      <h2>High Score: {{ flareView.data.score }}</h2>

      <div class="rounded-circle "
        style="width: 4rem; height: 4rem; background-color: black; color: white; display: flex; justify-content: center; align-items: center;">
        <p style="margin: 0;">bird🐤</p>
      </div>
    </div>
  </div> -->

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
  background: blue;
}

  
</style>

<script>
define(function(){
  return {
    props: ['flareView'],
    template: template,
    data: function() {
      return {amount: 0};
    },
    methods: {
      submit: function() {
        var self = this;
        var formData = {amount: parseInt(self.amount)};

        $flare.http.post('<% .Helpers.UrlForRoute "payment.received" %>', formData)
          .then(function(response){
            console.log(response);
            self.flareView.data.score = response.score;
          })
          .catch(function(error){
            console.log(error);
          });
      }
    }
  };
});

class Game {
 constructor(canvas, context){
  this.canvas = canvas;
  this.context = context
 } 
}
</script>
