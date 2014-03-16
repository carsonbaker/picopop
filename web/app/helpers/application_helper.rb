module ApplicationHelper

  def current_route_name
    Rails.application.routes.router.recognize(request) do |route, matches, param|
      return route.name + "_path"
    end
  end

  def external_css
    if Rails.env.development?
      stylesheet_link_tag "external"
    else
      capture do
        concat stylesheet_link_tag "//yui.yahooapis.com/pure/0.3.0/pure-min.css"
      end
    end
  end

  def external_js
    if Rails.env.development?
      javascript_include_tag "external"
    else
      capture do
        concat javascript_include_tag "//ajax.googleapis.com/ajax/libs/jquery/1.10.2/jquery.min.js"
        concat javascript_include_tag "//cdnjs.cloudflare.com/ajax/libs/backbone.js/1.0.0/backbone-min.js"
        concat javascript_include_tag "//cdnjs.cloudflare.com/ajax/libs/underscore.js/1.5.2/underscore-min.js"
      end
    end
  end

  def page_title(t, options = {})
    options.reverse_merge!({:show_host => true})
    @page_title = @page_title.nil? ? t : t + ' &mdash; ' + @page_title
    @full_page_title = @page_title unless options[:show_host]
  end

  def columned_table(n_columns, array, *table_class)
    i = 0
    content_tag(:table, :class => table_class) do
      table_html = ""
      for item in array
        table_html << "<tr>" if i % n_columns == 0
        table_html << content_tag(:td) do
          yield item
        end
        table_html << "</tr>" if i % n_columns == (n_columns - 1)
        i += 1
      end
      table_html.html_safe
    end
  end

  def present(object, klass=nil)
    klass ||= "#{object.class}Presenter".constantize
    presenter = klass.new(object, self)
    yield presenter if block_given?
    presenter
  end

  def one_line(&block)
    haml_concat capture_haml(&block).gsub("\n", '').gsub('\\n', "\n").gsub("\s\s", "").html_safe
  end

end
