package request

type Image2Image struct {
	InitImages        []string `json:"init_images,omitempty"`        // List of base64-encoded images to send as base/original
	ResizeMode        int      `json:"resize_mode,omitempty"`        // I don't know If I got it right or not. See: webui-api/img2img RESIZE_MODE helper package
	DenoisingStrength float64  `json:"denoising_strength,omitempty"` // Determines how little respect the algorithm should have for image's content. At 0, nothing will change, and at 1 you'll get an unrelated image.
	ImageCFGScale     float64  `json:"image_cfg_scale"`
	Mask              string   `json:"mask,omitempty"`            // base64-encoded mask image. What to put inside the masked area before processing it with Stable Diffusion.
	MaskBlur          int      `json:"mask_blur,omitempty"`       // How much to blur the mask before processing, in pixels.
	InpaintingFill    int      `json:"inpainting_fill,omitempty"` // I don't know If I got it right or not. See: webui-api/img2img INPAINT_MASK_CONENT helper package

	// Upscale masked region to target resolution, do inpainting, downscale back and paste into original image
	//
	// Original field was `InpaintFullRes` but since the default value is `true`. it's quite tricky to do this in GO
	//
	//  So I decided to reverse it. This set to true and "inpaint_full_res": false and vice versa
	DoNotInpaintFullRes bool `json:"inpaint_full_res,omitempty"`

	InpaintFullResPadding  int            `json:"inpaint_full_res_padding,omitempty"` // Amount of pixels to take as sample around the inpaint areas
	InpaintingMaskInvert   int            `json:"inpainting_mask_invert,omitempty"`   // I don't know If I got it right or not. See: webui-api/img2img INPAINT_MODE helper package
	InitialNoiseMultiplier int            `json:"initial_noise_multiplier,omitempty"` // I don't even see this on the UI itself. Please, if you know, tell me about it.
	Prompt                 string         `json:"prompt,omitempty"`
	NegativePrompt         string         `json:"negative_prompt,omitempty"`
	Styles                 []string       `json:"styles,omitempty"`
	Seed                   int            `json:"seed,omitempty"` // A value that determines the output of random number generator - if you create an image with same parameters and seed as another image, you'll get the same result
	Subseed                int            `json:"subseed,omitempty"`
	SubseedStrength        int            `json:"subseed_strength,omitempty"`
	SeedResizeFromH        int            `json:"seed_resize_from_h,omitempty"`
	SeedResizeFromW        int            `json:"seed_resize_from_w,omitempty"`
	SamplerName            string         `json:"sampler_name,omitempty"`  // Either SamplerName or SamplerIndex will be used.
	SamplerIndex           string         `json:"sampler_index,omitempty"` // Either SamplerName or SamplerIndex will be used.
	BatchSize              int            `json:"batch_size,omitempty"`    // How many do you want to simultaneously generate.
	BatchCount             int            `json:"n_iter,omitempty"`        // How many times do you want to generate.
	Steps                  int            `json:"steps,omitempty"`         // How many times to improve the generated image iteratively; higher values take longer; very low values can produce bad results
	CFGScale               float64        `json:"cfg_scale,omitempty"`     // Classifier Free Guidance Scale - how strongly the image should conform to prompt - lower values produce more creative results
	Width                  int            `json:"width,omitempty"`
	Height                 int            `json:"height,omitempty"`
	RestoreFaces           bool           `json:"restore_faces,omitempty"`
	Tiling                 bool           `json:"tiling,omitempty"`
	DoNotSaveSamples       bool           `json:"do_not_save_samples,omitempty"`
	DoNotSaveGrid          bool           `json:"do_not_save_grid,omitempty"`
	Eta                    float64        `json:"eta,omitempty"`
	SChurn                 float64        `json:"s_churn,omitempty"`
	STmax                  float64        `json:"s_tmax,omitempty"`
	STmin                  float64        `json:"s_tmin,omitempty"`
	SNoise                 float64        `json:"s_noise,omitempty"`
	OverrideSettings       map[string]any `json:"override_settings,omitempty"`

	// Original field was `OverrideSettingsRestoreAfterwards` but since the default value is `true`. it's quite tricky to do this in GO
	//
	//  So I decided to reverse it. This set to true and "override_settings_restore_afterwards": false and vice versa
	DoNotOverrideSettingsRestoreAfterwards bool `json:"override_settings_restore_afterwards"`

	ScriptName string   `json:"script_name,omitempty"`
	ScriptArgs []string `json:"script_args,omitempty"`

	IncludeInitImages bool `json:"include_init_images,omitempty"` // I don't even know what this is. But it has just a little impact on the result images

	// Original field was `SendImages` but since the default value is `true`. it's quite tricky to do this in GO
	//
	//  So I decided to reverse it. This set to true and "send_images": false and vice versa
	DoNotSendImages bool `json:"send_images"`

	SaveImages bool `json:"save_iamges,omitempty"`

	AlwaysOnScripts map[string]any `json:"alwayson_scripts,omitempty"`

	// If true, Will Decode Images after received response from API
	DecodeAfterResult bool `json:"-"`
}
