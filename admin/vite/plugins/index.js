
import vue from '@vitejs/plugin-vue'
import createAutoImport from './auto-import'
import createCompression from './compression'
import createSetupExtend from './setup-extend'
import createAutoComponents from './auto-components'
import CreateBuildAnalysis from './build-analysis'
import vueJsx from '@vitejs/plugin-vue-jsx'
export default function createVitePlugins(viteEnv, isBuild = false) {
  const vitePlugins = [vue()]
  vitePlugins.push(vueJsx())
  vitePlugins.push(createAutoImport())
  vitePlugins.push(createAutoComponents())
  vitePlugins.push(createSetupExtend())
  isBuild && vitePlugins.push(CreateBuildAnalysis())
  isBuild && vitePlugins.push(...createCompression(viteEnv))
  return vitePlugins
}
