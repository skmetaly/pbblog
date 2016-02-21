var gulp = require('gulp');
var uglify = require('gulp-uglify');
var concat = require('gulp-concat');
var cssmin = require('gulp-cssmin');

gulp.task('uglify', function() {
  gulp.src(['bower_components/bootstrap/dist/js/bootstrap.min.js'])
    .pipe(concat('main.js'))
    .pipe(uglify())
    .pipe(gulp.dest('public/assets/js'))

  gulp.src(['bower_components/bootstrap/dist/css/bootstrap.min.css'])
	.pipe(concat('main.min.css'))
	.pipe(cssmin())
	.pipe(gulp.dest('public/assets/css'));

  gulp.src(['bower_components/bootstrap/dist/fonts/*'])
    .pipe(gulp.dest('public/assets/fonts'))        
});

gulp.task('default', ['uglify'], function(){});
