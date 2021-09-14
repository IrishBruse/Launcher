<template>
    <div class="flex-grow bg-primary">
        <iframe v-if="finishedLoading == true" :src="props.CurrentGame.Url" class="w-full h-full" frameborder="0"> </iframe>
    </div>
</template>

<script setup>
import { onMounted, ref } from "@vue/runtime-core";

const props = defineProps({ CurrentGame: Object });

const finishedLoading = ref(false);

onMounted(() => {
    window.addEventListener("DownloadMetadataEvent", () => {
        finishedLoading.value = true;
    });
});
</script>
