const path = require('path');
var webpack = require('webpack');

module.exports = {
	entry: {
		bundle: './static/ts/src/index.ts',
		connect_school: './static/ts/src/standalone/connect_school.ts',
	},
	// devtool: 'inline-source-map',
	module: {
		rules: [
			{
				test: /\.tsx?$/,
				use: [
					{
						loader: 'ts-loader',
						options: {
							transpileOnly: true,
						},
					},
				],
				include: path.resolve(__dirname, 'static/ts/src'),
				exclude: /node_modules/,
			},
		],
	},
	resolve: {
		extensions: ['.tsx', '.ts', '.js'],
	},
	mode: 'development',
	output: {
		filename: '[name].js',
		path: path.resolve(__dirname, 'static/dist'),
	},
	plugins: [
		new webpack.ProvidePlugin({
			$: 'jquery',
			jQuery: 'jquery',
		}),
	],
};
