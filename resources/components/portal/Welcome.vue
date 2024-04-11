<template>

  <h1>Hello {{ first_name }} {{ last_name }}</h1> 

</template>

<style>
  /* styles here */
</style>

<script>
define(function(){
  var component = {
    template: template,
    data: function() {
      return {
        first_name: "Vince",
        last_name: "Parcon"
      };
    },
    mounted: function(){
      function Person(name){
        this.name = name
      }

      Person.prototype.greet = function (){
        alert("Hello, my name is " + this.name)
      }

      var p = new Person("Kyle")
    }
  };

  return component;
})
</script>
