import Vue from 'vue';
import feather from 'vue-icon';
import VueKonva from 'vue-konva';
import App from './App.vue';

Vue.config.productionTip = false;

Vue.use(VueKonva);
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
