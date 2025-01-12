<script lang="ts">
  import { GetListOfImages } from '../wailsjs/go/main/App.js'
  import {Modal} from "flowbite-svelte";
  import Zoom from "./components/Zoom.svelte";

  let name: string
  let images: string[] = []
  let zoomOpen = false
  let zoomImage: string = ""

  // Intersection Observer action
  function lazyLoad(element: HTMLImageElement) {
    const observer = new IntersectionObserver(
            (entries) => {
              entries.forEach(entry => {
                if (entry.isIntersecting) {
                  const img = entry.target as HTMLImageElement
                  img.src = img.dataset.src + "?width=400"
                  observer.unobserve(img)
                }
              })
            },
            {
              rootMargin: "3000px" // Start loading images when they're 50px from viewport
            }
    )

    observer.observe(element)

    return {
      destroy() {
        observer.unobserve(element)
      }
    }
  }

  function openZoom(image: string): void {
    zoomImage = image
    zoomOpen = true
  }

  function getListOfImages(): void {
    GetListOfImages().then(result => images = result)
  }
  getListOfImages()
</script>

<main class="bg-gray-800 min-h-screen p-8">
  <div class="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-4 xl:grid-cols-5 gap-4">
    {#each images as image}
      <button class="aspect-square">
        <img
                data-src={image}
                use:lazyLoad
                on:click={() => openZoom(image)}
                alt={image}
                class="w-full h-full object-contain hover:scale-105 transition-transform"
                loading="lazy"
        />
      </button>
    {/each}
  </div>
</main>
<Modal xs bind:open={zoomOpen} classDialog="w-full max-h-screen">
  <Zoom image={zoomImage} />
</Modal>