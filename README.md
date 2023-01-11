# Machbase Neo Documents

This repository is for the **Machbase Neo** documentation published at [machbase.github.io/machbase](https://machbase.github.io/machbase)

# Writer's environment

## Ubuntu

1. `apt-get install ruby-full build-essential zlib1g-dev`
2. add PATH
    
    ```
    export GEM_HOME="$HOME/gems"
    export PATH="$HOME/gems/bin:$PATH"
    ```

3. `gem install jekyll bundler`
4. `bundle install`
5. `bundle exec jekyll serve`
