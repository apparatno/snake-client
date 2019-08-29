import Vue from 'vue';
import feather from 'vue-icon';
import App from './App.vue';

Vue.config.productionTip = false;

Vue.use(feather, {
  name: 'v-icon',
  props: {
    baseClass: {
      type: String,
      default: 'v-icon',
    },
    classPrefix: {
      type: String,
      default: 'v-icon-',
    },
  },
});

new Vue({
  render: h => h(App),
}).$mount('#app');
