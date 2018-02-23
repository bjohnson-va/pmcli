package vstore


// Namespace calculates the environment-specific VStore namespace for your request and any response processing you may wish to do.
// This should be called on all namespace arguments passed to any VStore operations
func Namespace(env env, originalNamespace string) string {
	if env == Local {
		user, err := NewUserInfo(env, ""); if err != nil {
			panic(err)
		}
		return user.Namespace(originalNamespace)
	}
	return originalNamespace
}
