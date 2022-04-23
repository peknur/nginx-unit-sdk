package application

import "encoding/json"

type Type string

const (
	TypePHP      Type = "php"
	TypePython   Type = "python"
	TypePerl     Type = "perl"
	TypeRuby     Type = "ruby"
	TypeJava     Type = "java"
	TypeExternal Type = "external"
	TypeGo       Type = TypeExternal
)

type Config struct {
	// shared options between all application languages
	Type             Type              `json:"type,omitempty"`
	Limits           *Limits           `json:"limits,omitempty"`
	Processes        *Processes        `json:"processes,omitempty"`
	WorkingDirectory string            `json:"working_directory,omitempty"`
	User             string            `json:"user,omitempty"`
	Group            string            `json:"group,omitempty"`
	Environment      map[string]string `json:"environment,omitempty"`
	Isolation        *Isolation        `json:"isolation,omitempty"`

	// PHP and Python targets
	//
	// PHP: If you specify targets, there should be no root, index, or script defined at the application level.
	//
	// Python: If you specify targets, there should be no module or callable defined at the application level.
	// Moreover, you can’t combine WSGI and ASGI targets within a single app.
	Targets map[string]Target `json:"targets,omitempty"`

	// Perl, PHP and Ruby script
	//
	// Perl: PSGI script path.
	// PHP: Filename of a root-based PHP script that Unit uses to serve all requests to the app.
	// Ruby: Rack script pathname, including the `.ru` extension: `/www/rubyapp/script.ru`.
	Script string `json:"script,omitempty"`

	// Java and PHP options
	//
	// Java: Array of strings defining JVM runtime options.
	// Unit itself exposes the -Dnginx.unit.context.path option that defaults to /; use it to customize the context path.
	//
	// PHP: PHPoptions that defines the php.ini location and options.
	Options interface{} `json:"options,omitempty"`

	// Go and Javascript executable. Pathname of the application, absolute or relative to `working_directory`.
	Executable string `json:"executable,omitempty"`

	// Ruby hooks. Pathname of the .rb file defining the event hooks to be called during the app’s lifecycle.
	Hooks string `json:"hooks,omitempty"`

	// Perl, Java, Python and Ruby threads. Integer that sets the number of worker threads per app process.
	// When started, each app process creates a corresponding number of threads to handle requests.
	// The default value is `1`.
	Threads int `json:"threads,omitempty"`

	// Perl, Java, Python and Ruby thread stack size. Integer that defines the stack size of
	// a worker thread (in bytes, multiple of memory page size; the minimum value is usually architecture specific).
	// The default value is system dependent and can be set with `ulimit -s <SIZE_KB>`.
	ThreadStackSize int `json:"thread_stack_size,omitempty"`

	// Java webapp. Pathname of the application’s packaged or unpackaged .war file.
	Webapp string `json:"webapp,omitempty"`

	// Java classpath. Array of paths to your app’s required libraries (may list directories or .jar files).
	Classpath []string `json:"classpath,omitempty"`

	// Python protocol. Hint to tell Unit that the app uses a certain interface; can be asgi or wsgi.
	Protocol string `json:"protocol,omitempty"`

	// Python path. Array of strings that represent additional Python module lookup paths; these values are prepended to sys.path.
	Path []string `json:"path,omitempty"`
}

func (c *Config) UnmarshalJSON(data []byte) error {
	type C Config
	cfg := C{}
	if err := json.Unmarshal(data, &cfg); err != nil {
		return err
	}
	*c = Config(cfg)
	switch cfg.Type {
	case TypeJava:
		typeCastJavaOptions(c)
	case TypePHP:
		typeCastPHPOptions(c)
	}
	return nil
}

type PHPOptions struct {
	// Pathname of the php.ini file with PHP configuration directives.
	File string `json:"file,omitempty"`
	// Objects for extra directives. Values in admin are set in PHP_INI_SYSTEM mode,
	// so the app can’t alter them; user values are set in PHP_INI_USER mode and may be updated in runtime.
	User  map[string]string `json:"user,omitempty"`
	Admin map[string]string `json:"admin,omitempty"`
}

type Target struct {
	// PHP
	Root   string `json:"root,omitempty"`
	Script string `json:"script,omitempty"`
	Index  string `json:"index,omitempty"`

	// Python
	Module   string `json:"module,omitempty"`
	Callable string `json:"callable,omitempty"`
}

type Limits struct {
	Timeout  int `json:"timeout,omitempty"`
	Requests int `json:"requests,omitempty"`
}

type Processes struct {
	Max         int `json:"max,omitempty"`
	Spare       int `json:"spare,omitempty"`
	IdleTimeout int `json:"idle_timeout,omitempty"`
}

type Isolation struct {
	Namespaces *Namespaces `json:"namespaces,omitempty"`
	Uidmap     *[]UIDMap   `json:"uidmap,omitempty"`
	Gidmap     *[]GIDMap   `json:"gidmap,omitempty"`
	Rootfs     string      `json:"rootfs,omitempty"`
	Automount  *Automount  `json:"automount,omitempty"`
}

type Namespaces struct {
	Cgroup     bool `json:"cgroup"`
	Credential bool `json:"credential"`
	Mount      bool `json:"mount"`
	Network    bool `json:"network"`
	PID        bool `json:"pid"`
	Uname      bool `json:"uname"`
}

type UIDMap struct {
	Host      int `json:"host"`
	Container int `json:"container"`
	Size      int `json:"size"`
}
type GIDMap struct {
	Host      int `json:"host"`
	Container int `json:"container"`
	Size      int `json:"size"`
}

type Automount struct {
	LanguageDeps bool `json:"language_deps"`
	Procfs       bool `json:"procfs"`
	Tmpfs        bool `json:"tmpfs"`
}

func typeCastJavaOptions(c *Config) {
	if val, ok := c.Options.([]interface{}); ok {
		ops := make([]string, 0)
		for _, v := range val {
			ops = append(ops, v.(string))
		}
		c.Options = ops
	}
}

func typeCastPHPOptions(c *Config) {
	v, ok := c.Options.(map[string]interface{})
	if !ok || v == nil {
		return
	}
	ops := PHPOptions{}
	if val, ok := v["file"].(string); ok {
		ops.File = val
	}
	if val, ok := v["user"].(map[string]interface{}); ok {
		ops.User = make(map[string]string)
		for k, v := range val {
			ops.User[k] = v.(string)
		}
	}
	if val, ok := v["admin"].(map[string]interface{}); ok {
		ops.Admin = make(map[string]string)
		for k, v := range val {
			ops.Admin[k] = v.(string)
		}
	}
	c.Options = ops
}
