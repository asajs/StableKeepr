<script lang="ts">
  import { GetListOfImages } from '../wailsjs/go/main/App.js'

  let name: string
  let images: string[] = []

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

  function getListOfImages(): void {
    GetListOfImages().then(result => images = result)
  }
  getListOfImages()
</script>

<main class="bg-gray-800 min-h-screen p-8">
  <div class="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-4 xl:grid-cols-5 gap-4">
    {#each images as image}
      <div class="aspect-square">
        <img
                data-src={image}
                use:lazyLoad
                alt={image}
                class="w-full h-full object-contain hover:scale-105 transition-transform"
                loading="lazy"
        />
      </div>
    {/each}
  </div>
</main>