<template>
    <div>
        <h1>Welcome {{ flareView.data.name }}</h1>
        <p>Some other text</p>
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