<template>
    <div class="mx-auto p-2">
        <h1>Welcome {{ flareView.data.name }}</h1>

        <ul class="list-group list-group-horizontal mb-2" style="width: 90vw; max-width: 45rem;">
            <input type="text" class="form-control mr-2" placeholder="Username" aria-label="Username" aria-describedby="basic-addon1">
            <input type="text" class="form-control" placeholder="Username" aria-label="Username" aria-describedby="basic-addon1">
        </ul>

        <button type="button" class="btn btn-success mb-3">Sign In</button>

        <div class="rounded-lg p-4 mx-auto" style="width: 90vw; max-width: 45rem; height: 30rem; background-color: antiquewhite;">
           <h2>High Score: {{ flareView.data.score }}</h2>
           
            <div class="rounded-circle " style="width: 4rem; height: 4rem; background-color: black; color: white; display: flex; justify-content: center; align-items: center;">
                <p style="margin: 0;">bird🐤</p>
            </div>
        </div>

    </div>
</template>

<script>
define(function () {
    return {
        props: ['flareView'],
        template: template,
    }
})
</script>
