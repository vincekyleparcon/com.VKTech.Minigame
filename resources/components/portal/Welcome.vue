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

</script>
