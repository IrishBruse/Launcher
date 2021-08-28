<template>
    <img :src="src" class="modalImage noTextHighlight" @click="toggleModal" ref="sourceImage" />
    <teleport to="body">
        <div v-show="imageIsModal" class="modalImageDim" @click="closeModal">
            <img :src="src" class="modalImagePopup noTextHighlight" ref="modalImage" />
        </div>
    </teleport>
</template>

<script setup>
const { ref } = require("@vue/reactivity");
const { onMounted } = require("@vue/runtime-core");

const props = defineProps({
    src: String,
});

const imageIsModal = ref(false);

const sourceImage = ref();
const modalImage = ref();

const toggleModal = () => {
    imageIsModal.value = !imageIsModal.value;

    setTimeout(() => {
        if (imageIsModal.value) {
            if (modalImage.value.naturalHeight > modalImage.value.naturalWidth) {
                modalImage.value.classList.add("tall");
            } else {
                modalImage.value.classList.add("wide");
            }
        }
    }, 1);

    console.log(imageIsModal.value);
};

const closeModal = () => {
    modalImage.value.classList.remove("active");
    imageIsModal.value = false;

    modalImage.value.classList.remove("wide");
    modalImage.value.classList.remove("tall");
};
</script>

<style>
.modalImage {
    cursor: pointer;
}

.modalImagePopup {
    position: fixed;
    left: 50%;
    top: calc(50% + var(--navbarHeight) / 2);
    transform: translate(-50%, -50%);
    width: 50%;
    height: 50%;

    z-index: 10;

    transition: all 0.25s;
}

.modalImagePopup.wide {
    width: calc(95vw);
    height: auto;
}

.modalImagePopup.tall {
    width: auto;
    height: calc(95vh - var(--navbarHeight));
}

.modalImageDim {
    cursor: pointer;
    position: fixed;
    top: 0;
    left: 0;
    width: 100vw;
    height: 100vh;
    z-index: 5;
    background-color: #000000aa;
}
</style>