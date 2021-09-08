
<template>
    <div class="flex-grow bg-primary-light" ref="list">
        <button v-for="(project, i) in projectsData" :key="project.Name" class="gameSelect" @click="changeProject(i, project)">
            <img :src="project.IconURL" class="h-full bg-secondary transition-all duration-300 rounded-3xl group-hover:rounded-xl" />
            <h1 class="text-left flex-grow self-center px-2 overflow-hidden overflow-ellipsis">{{ project.Name }}</h1>
        </button>
    </div>
</template>

<script setup>
import { onMounted, ref } from "vue";

const emit = defineEmits(["projectChanged"]);

const projectsData = ref();
const list = ref();
const projects = ref([]);

const changeProject = (index, currentProject) => {
    for (let i = 0; i < list.value.children.length; i++) {
        var current = index == i;
        list.value.children[i].classList.toggle("current", current);
    }

    emit("projectChanged", currentProject.Name, currentProject.Versions);
};

function handleDownload() {
    projectsData.value = projectsMetadata;

    setTimeout(() => {
        changeProject(0, projectsData.value[0]);
    }, 1);
    document.removeEventListener("DownloadedProjectsMetadata", handleDownload);
}

onMounted(() => {
    document.addEventListener("DownloadedProjectsMetadata", handleDownload);
});
</script>

<style >
.gameSelect {
    @apply group transition duration-300 w-full h-14 relative flex p-2 cursor-pointer hover:bg-primary-hover;
}

.gameSelect.current > div {
    @apply scale-y-100;
}

.gameSelect.current > img {
    @apply left-1 rounded-xl;
}

.gameSelect.current > h1 {
    @apply font-bold;
}

.clipLoadingCircle {
    clip-path: polygon(0% 0%, 100% 0%, 100% 50%, 0% 50%);
}
</style>