#!/usr/bin/env ruby

def health_check
    begin
        File.read('/proc/1/cmdline')
        exit 0
    rescue
        exit 1
    end
end
  
# Call the function
health_check
