<template>
    <div class="flex-grow bg-primary">
        <iframe v-if="finishedLoading == true" :src="props.CurrentGame.Url" class="w-full h-full" frameborder="0"> </iframe>
        <div v-else class="flex justify-center h-full">
            <h1 class="justify-start self-center px-2 text-4xl">Loading</h1>
            <div class="w-16 h-16 self-center">
                <h2 class="absolute icons text-6xl text-tertiary clipLoadingCircle animate-spin">circle</h2>
            </div>
        </div>
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
