package dominantcolor

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"math/rand"
	"net/http"
	"time"

	"github.com/lucasb-eyer/go-colorful"
	"github.com/nfnt/resize"
)

func GetDominantColor(content []byte) (color.RGBA, error) {
	var img image.Image

	contentType := http.DetectContentType(content)

	if contentType == "image/png" {
		// Decode the image.
		newimg, err := png.Decode(bytes.NewReader(content))
		if err != nil {
			return color.RGBA{}, err
		}

		img = newimg
	} else {
		// Decode the image.
		newimg, err := jpeg.Decode(bytes.NewReader(content))
		if err != nil {
			return color.RGBA{}, err
		}

		img = newimg
	}

	// Resize the image to a smaller size to speed up processing.
	img = resize.Resize(50, 0, img, resize.Lanczos3)

	// Extract colors from the image.
	var colors []color.RGBA
	bounds := img.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			c, ok := colorful.MakeColor(img.At(x, y))

			if !ok {
				return color.RGBA{}, errors.New("cannot make color")
			}

			rgba := color.RGBA{
				uint8(c.R * 255),
				uint8(c.G * 255),
				uint8(c.B * 255),
				255,
			}
			colors = append(colors, rgba)
		}
	}

	// Use K-means to find the dominant color.
	clusters, err := kmeans(colors, 1)

	if err != nil {
		return color.RGBA{}, errors.New("cannot do kmeans algorithm")
	}

	colour := getDominant(clusters)

	return colour, nil
}

func getDominant(clusters []cluster) color.RGBA {
	var maxClusterIndex int
	maxPoints := 0

	for i, cluster := range clusters {
		if len(cluster.Points) > maxPoints {
			maxPoints = len(cluster.Points)
			maxClusterIndex = i
		}
	}

	dominantColor := clusters[maxClusterIndex].Center

	dominantColorRGBA := color.RGBA{
		R: uint8(dominantColor.R),
		G: uint8(dominantColor.G),
		B: uint8(dominantColor.B),
		A: 255,
	}

	return dominantColorRGBA
}

type poin struct {
	R, G, B float64
}

// cluster represents a cluster of points.
type cluster struct {
	Center poin
	Points []color.RGBA
}

func kmeans(colors []color.RGBA, k int) ([]cluster, error) {
	if k <= 0 || k > len(colors) {
		return nil, fmt.Errorf("invalid number of clusters")
	}

	// Convert color.RGBA to Point
	var points []poin
	for _, c := range colors {
		points = append(points, poin{R: float64(c.R), G: float64(c.G), B: float64(c.B)})
	}

	// Initialize clusters with random centroids.
	rand.Seed(time.Now().UnixNano())
	clusters := initializeClusters(points, k)

	// Run K-means iterations.
	for i := 0; i < 100; i++ { // You can adjust the number of iterations based on your needs.
		// Assign each point to the nearest cluster.
		assignPointsToClusters(points, clusters)

		// Update cluster centroids.
		if !updateCentroids(clusters) {
			break
		}
	}

	return clusters, nil
}

// initializeClusters initializes clusters with random centroids.
func initializeClusters(points []poin, k int) []cluster {
	clusters := make([]cluster, k)
	for i := 0; i < k; i++ {
		centroidIndex := rand.Intn(len(points))
		clusters[i].Center = points[centroidIndex]
	}
	return clusters
}

// assignPointsToClusters assigns each point to the nearest cluster.
func assignPointsToClusters(points []poin, clusters []cluster) {
	for i := range clusters {
		clusters[i].Points = nil
	}

	for _, point := range points {
		nearestClusterIndex := findNearestCluster(point, clusters)
		clusters[nearestClusterIndex].Points = append(clusters[nearestClusterIndex].Points, color.RGBA{
			R: uint8(point.R),
			G: uint8(point.G),
			B: uint8(point.B),
			A: 255,
		})
	}
}

// findNearestCluster finds the index of the nearest cluster for a given point.
func findNearestCluster(point poin, clusters []cluster) int {
	minDistance := distance(point, clusters[0].Center)
	nearestClusterIndex := 0

	for i := 1; i < len(clusters); i++ {
		d := distance(point, clusters[i].Center)
		if d < minDistance {
			minDistance = d
			nearestClusterIndex = i
		}
	}

	return nearestClusterIndex
}

// updateCentroids updates cluster centroids and returns true if any centroids changed.
func updateCentroids(clusters []cluster) bool {
	centroidsChanged := false

	for i := range clusters {
		oldCentroid := clusters[i].Center

		if len(clusters[i].Points) > 0 {
			clusters[i].Center = calculateMean(clusters[i].Points)
		}

		if oldCentroid != clusters[i].Center {
			centroidsChanged = true
		}
	}

	return centroidsChanged
}

// distance calculates the Euclidean distance between two points in the 3D color space.
func distance(p1, p2 poin) float64 {
	dR := p1.R - p2.R
	dG := p1.G - p2.G
	dB := p1.B - p2.B
	return dR*dR + dG*dG + dB*dB
}

// calculateMean calculates the mean point of a set of points.
func calculateMean(points []color.RGBA) poin {
	if len(points) == 0 {
		return poin{}
	}

	var sumR, sumG, sumB float64
	for _, p := range points {
		sumR += float64(p.R)
		sumG += float64(p.G)
		sumB += float64(p.B)
	}

	meanR := sumR / float64(len(points))
	meanG := sumG / float64(len(points))
	meanB := sumB / float64(len(points))

	return poin{meanR, meanG, meanB}
}
