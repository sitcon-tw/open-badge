###
 *
 *  Web Starter Kit
 *  Copyright 2014 Google Inc. All rights reserved.
 *
 *  Licensed under the Apache License, Version 2.0 (the "License")
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *      http:#www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License
 *
###

# Include Gulp & Tools We'll Use
gulp = require('gulp')
$ = require('gulp-load-plugins')()
del = require('del')
runSequence = require('run-sequence')
browserSync = require('browser-sync')
pagespeed = require('psi')
reload = browserSync.reload

AUTOPREFIXER_BROWSERS = [
  'ie >= 10',
  'ie_mob >= 10',
  'ff >= 30',
  'chrome >= 34',
  'safari >= 7',
  'opera >= 23',
  'ios >= 7',
  'android >= 4.4',
  'bb >= 10'
]

# Lint JavaScript
gulp.task 'jshint', () ->
  return gulp.src('app/scripts/**/*.js')
    .pipe(reload(stream: true, once: true))
    .pipe($.jshint())
    .pipe($.jshint.reporter('jshint-stylish'))
    .pipe($.if(!browserSync.active, $.jshint.reporter('fail')))

# Optimize Images
gulp.task 'images', () ->
  return gulp.src('app/images/**/*')
    .pipe($.cache($.imagemin({
      progressive: true,
      interlaced: true
    })))
    .pipe(gulp.dest('public/images'))
    .pipe($.size(title: 'images'))

# Copy All Files At The Root Level (app)
gulp.task 'copy', () ->
  return gulp.src(['app/*','!app/*.html'])
    .pipe(gulp.dest('public'))
    .pipe($.size(title: 'copy'))

# Copy Web Fonts To Dist
gulp.task 'fonts', () ->
  return gulp.src(['app/fonts/**'])
    .pipe(gulp.dest('public/fonts'))
    .pipe($.size(title: 'fonts'))

# Automatically Prefix CSS
gulp.task 'styles:css', () ->
  return gulp.src('app/styles/**/*.css')
    .pipe($.changed('app/styles'))
    .pipe($.autoprefixer(AUTOPREFIXER_BROWSERS))
    .pipe(gulp.dest('app/styles'))
    .pipe($.size(title: 'styles:css'))

# Compile Sass For Style Guide Components (app/styles/components)
gulp.task 'styles:components', () ->
  return gulp.src('app/styles/components/components.scss')
    .pipe($.rubySass({
      style: 'expanded',
      precision: 10,
      loadPath: ['app/styles/components']
      compass: true
    }))
    .on('error', console.error.bind(console))
    .pipe($.autoprefixer(AUTOPREFIXER_BROWSERS))
    .pipe(gulp.dest('app/styles/components'))
    .pipe($.size(title: 'styles:components'))

# Compile Any Other Sass Files You Added (app/styles)
gulp.task 'styles:scss', () ->
  return gulp.src(['app/styles/**/*.scss', '!app/styles/components/components.scss'])
    .pipe($.rubySass({
      style: 'expanded',
      precision: 10,
      loadPath: ['app/styles']
    }))
    .on('error', console.error.bind(console))
    .pipe($.autoprefixer(AUTOPREFIXER_BROWSERS))
    .pipe(gulp.dest('.tmp/styles'))
    .pipe($.size(title: 'styles:scss'))

# Output Final CSS Styles
gulp.task('styles', ['styles:components', 'styles:scss', 'styles:css'])

# Scan Your HTML For Assets & Optimize Them
gulp.task 'html', () ->
  return gulp.src('app/**/*.html')
    .pipe($.useref.assets(searchPath: '.tmp,app'))
    # Concatenate And Minify JavaScript
    .pipe($.if('*.js', $.uglify(preserveComments: 'some')))
    # Remove Any Unused CSS
    # Note: If not using the Style Guide, you can delete it from
    # the next line to only include styles your project uses.
    .pipe($.if('*.css', $.uncss({
      html: [
        'app/index.html',
        'app/styleguide/index.html'
      ],
      # CSS Selectors for UnCSS to ignore
      ignore: [
        '.navdrawer-container.open',
        /.app-bar.open/
      ]
    })))
    # Concatenate And Minify Styles
    .pipe($.if('*.css', $.csso()))
    .pipe($.useref.restore())
    .pipe($.useref())
    # Update Production Style Guide Paths
    .pipe($.replace('components/components.css', 'components/main.min.css'))
    # Minify Any HTML
    .pipe($.if('*.html', $.minifyHtml()))
    # Output Files
    .pipe(gulp.dest('public'))
    .pipe($.size(title: 'html'))

# Clean Output Directory
gulp.task('clean', del.bind(null, ['.tmp', 'public/*']))

# Watch Files For Changes & Reload
gulp.task 'serve', () ->
  browserSync({
    proxy: "localhost:8010"
    notify: true,
    server:{
      baseDir: ['.tmp', 'app']
    }
  })

  gulp.watch(['app/**/*.html'], reload)
  gulp.watch(['app/styles/**/*.scss'], ['styles:components', 'styles:scss'])
  gulp.watch(['.tmp,app/styles/**/*.css'], ['styles:css', reload])
  gulp.watch(['app/scripts/**/*.js'], ['jshint'])
  gulp.watch(['app/images/**/*'], reload)

# Build and serve the output from the dist build
gulp.task 'serve:dist', ['default'], () ->
  browserSync({
    notify: false,
    server:{
      baseDir: 'dist'
    }
  })

# Build Production Files, the Default Task
gulp.task 'default', ['clean'], (cb) ->
  runSequence('styles', ['jshint', 'html', 'images', 'fonts', 'copy'], cb)

# Run PageSpeed Insights
# Update `url` below to the public URL for your site
gulp.task 'pagespeed', pagespeed.bind(null,
  # By default, we use the PageSpeed Insights
  # free (no API key) tier. You can use a Google
  # Developer API key if you have one. See
  # http:#goo.gl/RkN0vE for info key: 'YOUR_API_KEY'
  url: 'http://localhsot:8010',
  strategy: 'mobile'
)

# Load custom tasks from the `tasks` directory
try
  require('require-dir')('tasks')
catch err