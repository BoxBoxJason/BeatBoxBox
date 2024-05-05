<template>
  <div class="home">
    <img alt="Vue logo" src="../assets/images/logo.png">
    <h1>Welcome to BeatBoxBox, a music platform to listen to your favourite songs ! </h1>
  </div>
  <div>
    <MusicSection title="Trending Musics" :musics="trendingMusics" />
    <MusicSection title="Recently Added Musics" :musics="recentlyAddedMusics " />
    <MusicSection title="Daily Recommendations" :musics="dailyRecommendations" />
  </div>
</template>

<script>
import MusicSection from '@/components/section/MusicSection.vue'

export default {
  name: 'HomeView',
  components: {
    MusicSection
  },
  data() {
    return{
      trendingMusics: [],
      recentlyAddedMusics: [],
      dailyRecommendations: []
    }
  }

}
</script>

<style>
.trendingMusics {
  background-color: burlywood;
} 
</style>
