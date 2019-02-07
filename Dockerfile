# Start from an empty image.
FROM scratch
# Copy the templates to the root of the container.
ADD template /template
# Copy the static binary to the root of the container.
ADD app /
# Execute it.
CMD ["/app"]
