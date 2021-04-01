import Vue from 'vue'
import { wrapFunctional } from './index'

const components = {
}

for (const name in components) {
  Vue.component(name, components[name])
  Vue.component('Lazy' + name, components[name])
}
