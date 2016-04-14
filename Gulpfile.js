'use strict';

var gulp = require('gulp');
var uglify = require('gulp-uglify');
var concat = require('gulp-concat');
var cssmin = require('gulp-cssmin');
var sass = require('gulp-sass');

var assetsPath = "resources/assets";

gulp.task('admin_uglify', function() {
  gulp.src([
    'bower_components/Materialize/dist/js/materialize.js'
    ])
    .pipe(concat('main.min.js'))
    .pipe(uglify())
    .pipe(gulp.dest('public/assets/js'))

});

gulp.task('admin_css',function(){
  gulp.src([
    'bower_components/Materialize/dist/css/materialize.min.css',
    assetsPath+'/css/admin/material-design.css',
    assetsPath+'/css/admin/*.css'
    ]
  )
  .pipe(concat('main.min.css'))
  .pipe(cssmin())
  .pipe(gulp.dest('public/assets/css'));
  });

gulp.task('admin_sass', function () {
  return gulp.src(assetsPath+'/scss/admin/**/*.scss')
    .pipe(sass().on('error', sass.logError))
    .pipe(gulp.dest(assetsPath+'/css/admin'));
});
 

gulp.task('fonts',function(){

  gulp.src(['bower_components/Materialize/dist/font/**/*'])
    .pipe(gulp.dest('public/assets/font'))          
})

gulp.task('assets:watch', function () {
  gulp.watch(assetsPath+'/scss/**/*.scss', ['admin_sass']);
  gulp.watch(assetsPath+'/css/admin/**/*.css', ['admin_css']);
});

gulp.task('admin', ['admin_uglify','admin_sass', 'admin_css','fonts'], function(){});

