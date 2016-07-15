module.exports = function(grunt) {
    var babel = require('rollup-plugin-babel');
    var nodeResolve = require('rollup-plugin-node-resolve');
    var commonjs = require('rollup-plugin-commonjs');
    var riot = require('rollup-plugin-riot');
    var uglify = require('rollup-plugin-uglify');
    var ruReplace = require('rollup-plugin-replace');
    
    grunt.initConfig({
        pkg: grunt.file.readJSON('package.json'),

        watch: {
            client: {
                files: ['client/*.js', 'client/components/*.tag'],
                tasks: ['build']
            }
        },

        rollup: {
            options: {
                entry: './client/app.js',

                plugins: [
                    riot(),

                    nodeResolve({
                        main: true,
                        jsnext: true,
                        browser: true
                    }),

                    ruReplace({
                        'process.env.NODE_ENV': JSON.stringify('production')
                    }),

                    commonjs(),

                    babel({
                        exclude: 'node_modules/**',
                        presets: ['es2015-rollup']
                    })//,

                    /*uglify({
                        wrap: true
                    })*/
                ]
            },

            files: {
                src: 'client/app.js',
                dest: 'assets/app.js'
            }
        }
    });

    grunt.loadNpmTasks('grunt-rollup');
    grunt.loadNpmTasks('grunt-contrib-watch');
    grunt.registerTask('build', ['rollup']);
};