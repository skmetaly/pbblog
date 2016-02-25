var gulp = require('gulp');
var uglify = require('gulp-uglify');
var concat = require('gulp-concat');
var cssmin = require('gulp-cssmin');

gulp.task('uglify', function() {
  gulp.src([
    'bower_components/Materialize/dist/js/materialize.js'
    ])
    .pipe(concat('main.min.js'))
    .pipe(uglify())
    .pipe(gulp.dest('public/assets/js'))

});

gulp.task('minCSS',function(){
  gulp.src([
    'bower_components/Materialize/dist/css/materialize.min.css'
    ]
  )
  .pipe(concat('main.min.css'))
  .pipe(cssmin())
  .pipe(gulp.dest('public/assets/css'));
  });

gulp.task('moveFonts',function(){

  gulp.src(['bower_components/Materialize/dist/font/*'])
    .pipe(gulp.dest('public/assets/font'))          
  })

gulp.task('default', ['uglify','minCSS','moveFonts'], function(){});

