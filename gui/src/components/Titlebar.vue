<template>
    <div class="h-16 bg-primary-dark flex justify-center">
        <h1 class="text-xl self-center text-center">{{ props.CurrentGame.Name }}</h1>
    </div>
</template>

<script setup>
const props = defineProps({ CurrentGame: Object });
</script>
