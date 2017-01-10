# config valid only for current version of Capistrano
lock '3.7.1'

set :application, 'champak'
set :repo_url, 'https://github.com/kapmahc/champak.git'

# Default branch is :master
# ask :branch, `git rev-parse --abbrev-ref HEAD`.chomp

# Default deploy_to directory is /var/www/my_app_name
# set :deploy_to, '/var/www/my_app_name'
set :deploy_to, -> { "/var/www/#{fetch :app_domain, 'localhost'}" }

# Default value for :format is :airbrussh.
# set :format, :airbrussh

# You can configure the Airbrussh format using :format_options.
# These are the defaults.
# set :format_options, command_output: true, log_file: 'log/capistrano.log', color: :auto, truncate: :auto

# Default value for :pty is false
# set :pty, true

# Default value for :linked_files is []
# append :linked_files, 'config/database.yml', 'config/secrets.yml'
append :linked_files, 'config/database.yml', '.rbenv-vars', 'vendor/assets/images/logo/spree_50.png', 'public/robots.txt', 'config/initializers/locale.rb'

# Default value for linked_dirs is []
# append :linked_dirs, 'log', 'tmp/pids', 'tmp/cache', 'tmp/sockets', 'public/system'
append :linked_dirs, 'log', 'tmp/pids', 'tmp/cache', 'tmp/sockets', 'tmp/sessions', 'public/system', 'public/attachments'

# Default value for default_env is {}
# set :default_env, { path: '/opt/ruby/bin:$PATH' }

# Default value for keep_releases is 5
# set :keep_releases, 5


# rbenv
set :rbenv_type, :user # or :system, depends on your rbenv setup
set :rbenv_ruby, File.read('.ruby-version').strip

# database
set :db_remote_clean, true
set :assets_dir, %w(public/attachments)
set :local_assets_dir, %w(public/attachments)
set :disallow_pushing, true
set :db_dump_dir, -> { File.join(current_path, 'db') }
set :compressor, :bzip2

# nginx
set :nginx_config_name, -> { fetch :app_domain }
set :nginx_server_name, -> { fetch :app_domain }
set :nginx_use_ssl, true
