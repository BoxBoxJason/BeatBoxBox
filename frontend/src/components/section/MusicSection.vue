<template>
    <div class="section">
      <h2>{{ title }}</h2>
      <div class="music-list">
        <MusicMiniature
          v-for="music in musics"
          :key="music.id"
          :title="music.title"
          :artist="music.artist"
          :album="music.album"
          :album_id="music.album_id"
          :path="music.path"
        />
      </div>
    </div>
  </template>
  
  <script>
  import MusicMiniature from '@/components/music/MusicMiniature.vue'
  
  export default {
    name: 'MusicSection',
    components: {
      MusicMiniature
    },
    props: {
      title: String,
      musics: Array
    }
  }
  </script>
  