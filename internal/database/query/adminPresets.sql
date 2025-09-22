-- name: CreateAdminPreset :exec
INSERT INTO Admin_Presets (preset_name, preset_value, slug) VALUES (?, ?, ?);

-- name: GetAdminPresetBySlug :one
SELECT id, preset_name, preset_value, slug FROM Admin_Presets WHERE slug = ?;

-- name: ListAdminPresets :many
SELECT id, preset_name, preset_value, slug FROM Admin_Presets;

