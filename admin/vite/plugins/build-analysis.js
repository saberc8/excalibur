
// 引入rollup-plugin-visualizer模块
import { visualizer } from 'rollup-plugin-visualizer'
// q: 分析的英文单词怎么写
// a: analysis
export default function CreateBuildAnalysis() {
	return visualizer({
		open: true, //注意这里要设置为true，否则无效
		filename: 'stats.html', //分析图生成的文件名
		gzipSize: true, // 收集 gzip 大小并将其显示
		brotliSize: true, // 收集 brotli 大小并将其显示
	})
}
