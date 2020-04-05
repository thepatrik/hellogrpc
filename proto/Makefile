.DEFAULT_GOAL := gen

PROTO_FILES = $(wildcard $(PROTO_DIR)/*.proto)
SERVICES = $(patsubst %.proto,%,$(notdir $(PROTO_FILES)))
PROTO_DIR = .
OUTPUT_DIR = build
PB_GO_SUFFIX = .pb.go
PB_JS_SUFFIX = _pb.js
PB_TS_SUFFIX = _pb.d.ts
MKDIR_P = mkdir -p
RMDIR = rm -rf

# make sure output dir is created
${OUTPUT_DIR}:
	${MKDIR_P} ${OUTPUT_DIR}

# default goal to generate protobuf code
.PHONY: gen
gen: _dirs _go _ts

# wipe clean
.PHONY: clean
clean: _dirs
	${RMDIR} ${OUTPUT_DIR}
	${RMDIR} node_modules

# create necessary directories
.PHONY: _dirs
_dirs: ${OUTPUT_DIR}

# needed for TypeScript and JavaScript protobuf code generation
node_modules: package.json package-lock.json
	npm install --no-save

# generates go protobuf code
.PHONY: _go
_go: $(addprefix $(OUTPUT_DIR)/,$(addsuffix $(PB_GO_SUFFIX),$(SERVICES)))
$(OUTPUT_DIR)/%$(PB_GO_SUFFIX): $(PROTO_DIR)/%.proto
	protoc \
	-I $(PROTO_DIR) \
	--go_out=plugins=grpc:$(OUTPUT_DIR) \
	$<

# generates protobuf code for JavaScript
.PHONY: _js
_js: node_modules $(addprefix $(OUTPUT_DIR)/,$(addsuffix $(PB_JS_SUFFIX),$(SERVICES)))
$(OUTPUT_DIR)/%$(PB_JS_SUFFIX): $(PROTO_DIR)/%.proto
	node_modules/.bin/grpc_tools_node_protoc \
	--js_out=import_style=commonjs,binary:${OUTPUT_DIR} \
	--grpc_out=${OUTPUT_DIR} \
	--plugin=protoc-gen-grpc=node_modules/.bin/grpc_tools_node_protoc_plugin \
	-I $(PROTO_DIR) \
	$<

# generates TypeScript definition protobuf code
.PHONY: _ts
_ts: node_modules _js $(addprefix $(OUTPUT_DIR)/,$(addsuffix $(PB_TS_SUFFIX),$(SERVICES)))
$(OUTPUT_DIR)/%$(PB_TS_SUFFIX): $(PROTO_DIR)/%.proto
	protoc \
	--plugin=protoc-gen-ts=./node_modules/.bin/protoc-gen-ts \
	--ts_out=${OUTPUT_DIR} \
	-I $(PROTO_DIR) \
	$<
